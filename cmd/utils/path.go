package utils

import "os"

func GetPath() string {
	pathEnv, exists := os.LookupEnv("PATH")
	if exists {
		return pathEnv
	}
	return ""
}
