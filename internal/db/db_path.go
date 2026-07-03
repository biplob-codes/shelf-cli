package db

import (
	"fmt"
	"os"
	"path/filepath"
)

func getDataSource() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("getting config directory: %w", err)
	}
	dir := filepath.Join(configDir, "shelf")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("creating shelf directory: %w", err)
	}
	return filepath.Join(dir, "shelf.db"), nil
}
