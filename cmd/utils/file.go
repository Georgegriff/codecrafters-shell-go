package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FindFile(dir string, fileQuery string) (string, bool) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == fileQuery && !file.IsDir() {
			filePath := filepath.Join(dir, file.Name())
			return filePath, true
		}
	}

	return "", false
}

func FindFileInPath(fileQuery string) (string, bool) {
	path := GetPath()

	pathDirectories := strings.Split(path, ":")

	for _, dir := range pathDirectories {
		file, found := FindFile(dir, fileQuery)
		if found {
			return file, true
		}
	}
	return "", false
}
