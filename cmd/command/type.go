package command

import (
	"fmt"
	"os"
)

type Type struct {
}

func (e Type) Run(args []string) {
	lookUpCommand := ""
	if len(args) >= 1 {
		lookUpCommand = args[0]
	}
	command, err := GetCommand(CommandType(lookUpCommand))
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "%s is a shell builtin", command)
}

func (e Type) String() string {
	return string(TYPE)
}
