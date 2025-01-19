package command

import (
	"fmt"
	"os"
	"strings"
)

type Echo struct {
}

func (e Echo) CommandType() CommandType {
	return ECHO
}

func (e Echo) String() string {
	return string(e.CommandType())
}

func (e Echo) Run(args []string) {
	fmt.Fprint(os.Stdout, strings.Join(args, " "), "\n")
}
