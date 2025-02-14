package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const DbName = "/.gatorconfig.json"

type Config struct {
	DbUrl  string `json:"db_url"`
	DbUser string `json:"current_user_name"`
}

func Read() (*Config, error) {
	// Example: Read from a JSON file or environment variables
	cfg := &Config{
		DbUrl: "postgres://grahamhill:@localhost:5432/gator?sslmode=disable",
	}
	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, DbName), nil
}

func (c *Config) SetUser(userName string) error {
	c.DbUser = userName
	return Write(c)
}

func Write(c *Config) error {
	// Get the path to the config file
	configPath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("failed to get config path: %w", err)
	}

	// Convert the config to JSON
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Write the config file
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
