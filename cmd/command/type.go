package command

import (
	"fmt"
	"os"
)

type Type struct {
	BaseCommand
}

func (t Type) CommandType() CommandType {
	return TYPE
}

func (t Type) Run(args []string) {
	lookUpCommand := ""
	if len(args) >= 1 {
		lookUpCommand = args[0]
	}
	command, err := GetCommand(CommandType(lookUpCommand))
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return
	}
	userDefined, isUserDefined := command.(UserDefined)
	if isUserDefined {
		fmt.Fprintf(os.Stdout, "%s is %s", userDefined.Name, userDefined.Path)
	} else {

		fmt.Fprintf(os.Stdout, "%s is a shell builtin", command)
	}

}
