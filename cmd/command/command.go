package command

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Command interface {
	Run(args []string)
	String() string
	CommandType() CommandType
}

func ExecuteCommandInput(commandInput string) {
	commandType, arguments := parseArgs(commandInput)

	command, err := GetCommand(CommandType(commandType))
	if err != nil {
		notFoundError := fmt.Errorf("%s: command not found", strings.Join(append([]string{string(commandType)}, arguments...), " "))
		fmt.Fprint(os.Stdout, notFoundError, "\n")
		return
	}
	command.Run(arguments)

}

// Horrible had to borrow this
func parseArgs(input string) (string, []string) {

	var args []string
	command, argstr, _ := strings.Cut(input, " ")
	if strings.Contains(input, "\"") {
		re := regexp.MustCompile("\"(.*?)\"")
		args = re.FindAllString(input, -1)
		for i := range args {
			args[i] = strings.Trim(args[i], "\"")
		}
	} else if strings.Contains(input, "'") {
		re := regexp.MustCompile("'(.*?)'")
		args = re.FindAllString(input, -1)
		for i := range args {
			args[i] = strings.Trim(args[i], "'")
		}
	} else {
		if strings.Contains(argstr, "\\") {
			re := regexp.MustCompile(`[^\\] +`)
			args = re.Split(argstr, -1)
			for i := range args {
				args[i] = strings.ReplaceAll(args[i], "\\", "")
			}
		} else {
			args = strings.Fields(argstr)
		}
	}

	return command, args
}
