package command

import (
	"fmt"
	"os"
	"strings"
	"unicode"
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

func parseArgs(s string) (string, []string) {
	var result []string
	var current []rune
	var quote rune
	var nestedQuote rune
	escaped := false

	for i, r := range s {
		switch {
		case escaped:
			current = append(current, r)
			escaped = false
		case r == '\\':
			if nestedQuote != '\'' && quote != '\'' {
				if quote == 0 || (quote != 0 && (s[i+1] == '"' || s[i+1] == '\\' || s[i+1] == '$')) {
					escaped = true
				} else {
					current = append(current, r)
				}
			} else {
				current = append(current, r)
			}
		case quote != 0:
			if r == quote {
				quote = 0
			} else {
				if r == '"' || r == '\'' {
					if nestedQuote == r {
						nestedQuote = 0
					} else {
						nestedQuote = r
					}
				}
				current = append(current, r)
			}
		case r == '"' || r == '\'':
			quote = r
		case unicode.IsSpace(r):
			if len(current) > 0 {
				result = append(result, string(current))
				current = nil
			}
		default:
			current = append(current, r)
		}
	}

	if len(current) > 0 {
		result = append(result, string(current))
	}

	return result[0], result[1:]
}
