package util

import (
	"os"
	"path/filepath"
)

func CreateDir(relativePath string) (absolutePath string, err error) {
	baseDir, _ := os.Getwd()
	absolutePath = filepath.Join(baseDir, relativePath)
	if _, err = os.Stat(absolutePath); os.IsNotExist(err) {
		err = os.Mkdir(absolutePath, os.ModePerm)
	}

	return
}
