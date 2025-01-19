package command

import (
	"fmt"
	"os"
	"strings"
)

type Command interface {
	Run(args []string)
	String() string
	CommandType() CommandType
}

func ExecuteCommandInput(commandInput string) {
	commandParts := strings.Split(commandInput, " ")

	commandType, arguments := commandParts[0], commandParts[1:]
	command, err := GetCommand(CommandType(commandType))
	if err != nil {
		notFoundError := fmt.Errorf("%s: command not found", strings.Join(append([]string{string(commandType)}, arguments...), " "))
		fmt.Fprint(os.Stdout, notFoundError)
		return
	}
	command.Run(arguments)

}
