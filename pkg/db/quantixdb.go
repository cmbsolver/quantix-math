package db

import (
	"errors"
	"fmt"
	"os"
	"quantix-math/pkg/config"
	"quantix-math/pkg/db/tables"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %v", err)
	}

	adminStr := cfg.AdminConnectionString

	adminConn, err := pgx.Connect(context.Background(), adminStr)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		if err != nil {
			return nil, err
		}
		//os.Exit(1)
	}
	defer func(adminConn *pgx.Conn, ctx context.Context) {
		err := adminConn.Close(ctx)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "Error closing database connection: %v\n", err)
			if err != nil {
				return
			}
		}
	}(adminConn, context.Background())

	// Create the database if it does not exist
	createDatabaseSQL := `CREATE DATABASE quantixdb;`
	_, err = adminConn.Exec(context.Background(), createDatabaseSQL)
	if err != nil {
		// Check if the error is "duplicate_database" (42P04)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "42P04" {
			fmt.Printf("Database already exists, continuing...\n")
		} else {
			return nil, err
		}
	}

	// Now we need to put in our migrations.
	conn, err := InitConnection()
	if err != nil {
		println(err.Error())
		return nil, err
	}

	// Migrate the schemas
	dbCreateError := conn.AutoMigrate(&tables.DictionaryWord{})
	if dbCreateError != nil {
		println(dbCreateError.Error())
		fmt.Printf("Error creating table: %v\n", dbCreateError)
	}

	return conn, nil
}

func InitConnection() (*gorm.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %v", err)
	}

	connStr := cfg.GeneralConnectionString

	dsn := connStr
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CloseConnection closes the database connection
func CloseConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error getting database instance: %v", err)
	}
	return sqlDB.Close()
}
