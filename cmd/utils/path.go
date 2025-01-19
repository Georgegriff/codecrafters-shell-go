package utils

import (
	"log"
	"os"
)

func GetPath() string {
	pathEnv, exists := os.LookupEnv("PATH")
	if exists {
		return pathEnv
	}
	return ""
}

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}
	return homeDir
}
