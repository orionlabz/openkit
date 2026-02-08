package platform

import (
	"os"
	"path/filepath"
)

func OpenKitHome() (string, error) {
	if v := os.Getenv("OPENKIT_HOME"); v != "" {
		return v, nil
	}
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userHome, ".openkit"), nil
}

func OpenKitBinDir() (string, error) {
	if v := os.Getenv("OPENKIT_INSTALL_DIR"); v != "" {
		return v, nil
	}
	home, err := OpenKitHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "bin"), nil
}

func OpenKitStatePath() (string, error) {
	home, err := OpenKitHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "state.json"), nil
}
