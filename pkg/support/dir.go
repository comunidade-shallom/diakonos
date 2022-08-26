package support

import (
	"os"
	"path/filepath"
)

func GetBinDirPath() string {
	execPath, _ := os.Executable()

	return filepath.Dir(execPath)
}

func EnsureDir(dirName string) error {
	if _, err := os.Stat(dirName); err != nil {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
