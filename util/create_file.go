package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFile(dirPath string, fileName string, extension string) (*os.File, error) {
	if fileName == "" {
		return nil, fmt.Errorf("empty file name")
	}

	if extension == "" {
		return nil, fmt.Errorf("empty file extension")
	}

	filePath := filepath.Join(dirPath, fmt.Sprintf("%v.%v", fileName, extension))

	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}
