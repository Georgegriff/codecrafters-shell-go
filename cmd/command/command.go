package command

import (
	"fmt"
	"os"
	"strings"
)

type Command interface {
	Run(args []string)
}

type CommandType string

const (
	EXIT CommandType = "exit"
	ECHO CommandType = "echo"
)

func ExecuteCommandInput(commandInput string) {
	commandParts := strings.Split(commandInput, " ")

	commandType, arguments := commandParts[0], commandParts[1:]
	command, err := getCommand(CommandType(commandType), arguments)
	if err != nil {
		fmt.Fprint(os.Stdout, err.Error())
		return
	}
	command.Run(arguments)

}

func getCommand(commandType CommandType, arguments []string) (Command, error) {
	switch commandType {
	case EXIT:
		return Exit{}, nil
	case ECHO:
		return Echo{}, nil
	default:
		return nil, fmt.Errorf("%s: command not found", strings.Join(append([]string{string(commandType)}, arguments...), " "))
	}
}
