package config

import (
	"encoding/json"
	"os"
)

// Defines a config struct based on gatorconfig.json
type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Helper function to Read Configuration File
func Read(filepath string) (Config, error) {
	var configResult Config
	content, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(content, &configResult)
	if err != nil {
		return Config{}, err
	}
	return configResult, nil
}

// Helper function to change the username in configuration file
func (c *Config) SetUser(username string, filepath string) error {
	c.CurrentUserName = username
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, data, 0644) // 0644 used to define read-write privileges to owner and read only privilege to others
	if err != nil {
		return err
	}
	return nil

}
