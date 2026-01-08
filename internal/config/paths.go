package config

import (
	"os"
	"path/filepath"
)

func DataDir(appName string) (string, error) {
	userDataLocation, err := os.UserConfigDir()

	if err != nil {
		return "", err
	}

	databaseLocation := filepath.Join(userDataLocation, appName)

	if err := os.MkdirAll(databaseLocation, 0o755); err != nil {
		return "", err
	}

	return databaseLocation, nil
}
