package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	if path == "" {
		path = utils.GetHomeDir()
	}
	if strings.HasPrefix(path, "~") {
		after, found := strings.CutPrefix(path, "~")
		if found {
			path = filepath.Join(utils.GetHomeDir(), after)
		}
	}
	if !utils.DirExists(path) {
		fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", path)
		return
	}

	err := os.Chdir(path)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}

}
