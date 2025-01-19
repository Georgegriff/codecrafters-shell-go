package command

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Cat struct {
}

func (c Cat) CommandType() CommandType {
	return CAT
}

func (c Cat) String() string {
	return string(c.CommandType())
}

func (c Cat) Run(args []string) {
	if len(args) < 1 {
		// TODO handle read from stdin, pipe to out and kill cmd
	}
	for _, path := range args {
		info, err := os.Stat(path)

		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stdout, "cat: %s: No such file or directory\n", path)
		} else if info.IsDir() {
			fmt.Fprintf(os.Stdout, "cat: %s: Is a directory\n", path)
		} else {
			printFile(path)
		}

	}
}

func printFile(path string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal("error opening file")
		return
	}
	defer file.Close()
	_, err = io.Copy(os.Stdout, file)

	if err != nil {
		log.Fatalf("error printing file to stdout: %e", err)
	}
}
