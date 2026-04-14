package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func Read() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not find home directory:", err)
	}
	filePath := getConfigFilepath(homeDir)

	contents, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("File could not be found:", err)
	}
	var c Config
	err = json.Unmarshal(contents, &c)
	if err != nil {
		log.Fatal("Error decoding json:", err)
	}

	return c
}

func getConfigFilepath(homeDir string) string {
	filepath := filepath.Join(homeDir, configFileName)
	return filepath
}
