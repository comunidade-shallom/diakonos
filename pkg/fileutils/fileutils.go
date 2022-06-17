package fileutils

import (
	"errors"
	"os"
	"path/filepath"
)

func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func GetRelative(fileName string) string {
	pwd, _ := os.Getwd()

	r, _ := filepath.Rel(pwd, fileName)

	return r
}
