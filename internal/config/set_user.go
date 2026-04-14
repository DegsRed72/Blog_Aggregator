package config

import (
	"encoding/json"
	"log"
	"os"
)

func (cfg *Config) SetUser(username string) {
	cfg.CurrentUserName = username
	err := write(*cfg)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}
}

func write(cfg Config) error {
	jsonData, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}
	homeDir, err := os.UserHomeDir()
	filepath := getConfigFilepath(homeDir)
	os.WriteFile(filepath, jsonData, 0644)
	return nil
}
