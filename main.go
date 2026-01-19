package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ValidateJSON validates if a file contains valid JSON
func ValidateJSON(filePath string) error {
	// Read the file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Try to unmarshal as generic JSON
	var js interface{}
	if err := json.Unmarshal(data, &js); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	return nil
}

// ValidateConfigFolder validates all JSON files in the config folder
func ValidateConfigFolder(configDir string) ([]string, []string) {
	var validFiles []string
	var invalidFiles []string

	// Read all files in the config directory
	files, err := ioutil.ReadDir(configDir)
	if err != nil {
		log.Printf("Error reading config directory: %v", err)
		return validFiles, invalidFiles
	}

	// Validate each JSON file
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Only process .json files
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		filePath := filepath.Join(configDir, file.Name())
		if err := ValidateJSON(filePath); err != nil {
			log.Printf("❌ %s: INVALID - %v", file.Name(), err)
			invalidFiles = append(invalidFiles, file.Name())
		} else {
			log.Printf("✅ %s: VALID", file.Name())
			validFiles = append(validFiles, file.Name())
		}
	}

	return validFiles, invalidFiles
}

func main() {
	// Get config directory from args or use default
	configDir := "config"
	if len(os.Args) > 1 {
		configDir = os.Args[1]
	}

	// Check if config directory exists
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		log.Fatalf("Config directory '%s' does not exist", configDir)
	}

	fmt.Println("===========================================")
	fmt.Println("JSON Configuration File Validator")
	fmt.Println("===========================================")
	fmt.Printf("Validating JSON files in: %s\n", configDir)
	fmt.Println("-------------------------------------------")

	validFiles, invalidFiles := ValidateConfigFolder(configDir)

	fmt.Println("===========================================")
	fmt.Printf("Summary: %d valid, %d invalid\n", len(validFiles), len(invalidFiles))
	fmt.Println("===========================================")

	// Exit with error code if any files are invalid
	if len(invalidFiles) > 0 {
		os.Exit(1)
	}
}
