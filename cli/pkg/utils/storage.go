package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetStoragePath(appName string) (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	var storagePath string

	switch runtime.GOOS {
	case "darwin": // macOS
		storagePath = filepath.Join(homeDir, "Library", "Application Support", appName)
	case "linux": // Linux
		storagePath = filepath.Join(homeDir, ".config", appName)
	case "windows": // Windows
		appData := os.Getenv("APPDATA")
		if appData == "" {
			return "", fmt.Errorf("APPDATA environment variable is not set")
		}
		storagePath = filepath.Join(appData, appName)
	default:
		return "", fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	// Ensure the directory exists
	if err := os.MkdirAll(storagePath, 0755); err != nil {
		return "", fmt.Errorf("failed to create storage directory: %w", err)
	}

	return storagePath, nil
}
