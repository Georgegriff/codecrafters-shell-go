package command

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/utils"
)

type Cd struct {
}

func (c Cd) CommandType() CommandType {
	return CD
}

func (c Cd) String() string {
	return string(c.CommandType())
}

func (c Cd) Run(args []string) {
	path := ""
	if len(args) >= 1 {
		path = args[0]
	}
	if path != "" {
		if !utils.DirExists(path) {
			fmt.Fprintf(os.Stdout, "cd: %s: Not such file or directory\n", path)
			return
		}

		err := os.Chdir(path)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
		return
	}
	// todo handle "" go to home

}
