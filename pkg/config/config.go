package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// AppConfig represents the application configuration.
type AppConfig struct {
	ExistingHash            string `json:"existing_hash"`
	AdminConnectionString   string `json:"admin_connection_string"`
	GeneralConnectionString string `json:"general_connection_string"`
}

// getConfigFilePath returns the path to the configuration file.
func getConfigFilePath() (string, error) {
	return filepath.Join("./settings", "appsettings.json"), nil
}

// LoadConfig reads the configuration file and returns the AppConfig struct.
func LoadConfig() (*AppConfig, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	// Check if the config file exists
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// If the config file does not exist, create a default config
		err = CreateDefaultConfig()
		if err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config AppConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// CreateDefaultConfig creates a default configuration file.
func CreateDefaultConfig() error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Remove the existing configuration file if it exists
	if _, err := os.Stat(configFilePath); err == nil {
		err = os.Remove(configFilePath)
		if err != nil {
			return err
		}
	}

	configDir := filepath.Dir(configFilePath)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			return err
		}
	}

	defaultConfig := AppConfig{
		ExistingHash:            "36367763ab73783c7af284446c59466b4cd653239a311cb7116d4618dee09a8425893dc7500b464fdaf1672d7bef5e891c6e2274568926a49fb4f45132c2a8b4",
		AdminConnectionString:   "postgres://postgres:quantixpw@localhost:5432/postgres",
		GeneralConnectionString: "host=localhost user=postgres password=quantixpw dbname=quantixdb port=5432 sslmode=disable TimeZone=America/Chicago search_path=public",
	}

	data, err := json.MarshalIndent(defaultConfig, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, data, 0644)
	if err != nil {
		return err
	}

	// Update NumWorkers to the number of CPU cores
	err = UpdateConfig("NumWorkers", runtime.NumCPU()*2)
	if err != nil {
		return err
	}

	return nil
}

// UpdateConfig updates a specific field in the configuration file and saves it.
func UpdateConfig(key string, value interface{}) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	var config AppConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	switch key {
	case "ExistingHash":
		config.ExistingHash = value.(string)
	case "AdminConnectionString":
		config.AdminConnectionString = value.(string)
	case "GeneralConnectionString":
		config.GeneralConnectionString = value.(string)
	default:
		return fmt.Errorf("unknown configuration key: %s", key)
	}

	data, err = json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(configFilePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
