package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-shellwords"
)

type Command interface {
	Run(args []string)
	String() string
	CommandType() CommandType
}

func ExecuteCommandInput(commandInput string) {
	commandParts := parseArgs(commandInput)

	commandType, arguments := commandParts[0], commandParts[1:]
	command, err := GetCommand(CommandType(commandType))
	if err != nil {
		notFoundError := fmt.Errorf("%s: command not found", strings.Join(append([]string{string(commandType)}, arguments...), " "))
		fmt.Fprint(os.Stdout, notFoundError, "\n")
		return
	}
	command.Run(arguments)

}

func parseArgs(input string) []string {

	p := shellwords.NewParser()
	p.ParseBacktick = true

	args, _ := p.Parse(input)

	return args
}
