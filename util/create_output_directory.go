package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateOutputDirectory(folderName string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %w", err)
	}

	dirPath := filepath.Join(cwd, "output", fmt.Sprintf("%v", folderName))
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	return dirPath, nil
}
