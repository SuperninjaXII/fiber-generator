package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateDir creates directories based on the given file path.
func CreateDir(filePath string) error {
	if err := os.MkdirAll(filepath.Join(filePath), 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}
	return nil
}

// CreateFile creates a file at the given path with the provided content.
func CreateFile(filePath, content string) error {
	// Create the directories if they don't exist
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("error creating directories: %w", err)
	}

	// Write the content to the file
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

