package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const DbName = "gatorconfig.json"

type Config struct {
	DbUrl  string `json:"db_url"`
	DbUser string `json:"current_user_name"`
}

func Read() (*Config, error) {
	dbPath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(dbPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func getConfigFilePath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Navigate to the parent directory
	parentDir := filepath.Dir(currentDir)

	// Construct the full path to the config file
	return filepath.Join(parentDir, DbName), nil
}

func (c *Config) SetUser(userName string) error {
	c.DbUser = userName
	return write(*c)
}

func write(c Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
