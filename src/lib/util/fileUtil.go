package util

import (
	"os"
	"path/filepath"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDirByPath(path string) bool {
	err := os.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		return false
	}
	return true
}

func CreateDirByDirPath(dirPath string) bool {
	err := os.MkdirAll(dirPath, 0777)
	if err != nil {
		return false
	}
	return true
}
