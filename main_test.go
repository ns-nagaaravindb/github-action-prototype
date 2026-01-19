package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestValidateJSON(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		wantError bool
	}{
		{
			name:      "valid JSON object",
			content:   `{"key": "value", "number": 123}`,
			wantError: false,
		},
		{
			name:      "valid JSON array",
			content:   `[1, 2, 3, "test"]`,
			wantError: false,
		},
		{
			name:      "invalid JSON - missing quote",
			content:   `{"key: "value"}`,
			wantError: true,
		},
		{
			name:      "invalid JSON - trailing comma",
			content:   `{"key": "value",}`,
			wantError: true,
		},
		{
			name:      "empty JSON object",
			content:   `{}`,
			wantError: false,
		},
		{
			name:      "empty JSON array",
			content:   `[]`,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file
			tmpfile, err := ioutil.TempFile("", "test_*.json")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			// Write test content
			if _, err := tmpfile.Write([]byte(tt.content)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			// Test validation
			err = ValidateJSON(tmpfile.Name())
			if (err != nil) != tt.wantError {
				t.Errorf("ValidateJSON() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestValidateConfigFolder(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := ioutil.TempDir("", "test_config_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test files
	testFiles := map[string]string{
		"valid1.json":  `{"test": "data"}`,
		"valid2.json":  `{"another": "file", "count": 42}`,
		"invalid.json": `{invalid json}`,
		"notjson.txt":  `this is not a json file`,
		"empty.json":   `{}`,
	}

	for filename, content := range testFiles {
		filePath := filepath.Join(tmpDir, filename)
		if err := ioutil.WriteFile(filePath, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Test the validation
	validFiles, invalidFiles := ValidateConfigFolder(tmpDir)

	// Check results
	if len(validFiles) != 3 {
		t.Errorf("Expected 3 valid files, got %d", len(validFiles))
	}

	if len(invalidFiles) != 1 {
		t.Errorf("Expected 1 invalid file, got %d", len(invalidFiles))
	}

	// Verify that txt file was skipped
	for _, file := range validFiles {
		if file == "notjson.txt" {
			t.Error("Non-JSON file should not be validated")
		}
	}
}
