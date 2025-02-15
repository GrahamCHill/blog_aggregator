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
	// Get the path to the config file
	configPath, err := getConfigFilePath()
	if err != nil {
		return nil, fmt.Errorf("failed to get config path: %w", err)
	}

	// Check if the config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// If the file doesn't exist, return a default config
		return &Config{
			DbUrl: "postgres://grahamhill:@localhost:5432/gator?sslmode=disable",
		}, nil
	}

	// Read the config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse the JSON into a Config struct
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
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
