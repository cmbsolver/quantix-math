package db

import (
	"fmt"
	"quantix-math/pkg/config"
	"quantix-math/pkg/db/tables"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
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

	dbPath := cfg.DBPath
	if dbPath == "" {
		dbPath = "quantix.db"
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
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
