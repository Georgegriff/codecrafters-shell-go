package command

import (
	"fmt"
	"os"
	"strings"
)

type Echo struct {
	BaseCommand
}

func (e Echo) CommandType() CommandType {
	return ECHO
}

func (e Echo) Run(args []string) {
	fmt.Fprint(os.Stdout, strings.Join(args, " "))
}
