package command

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/cmd/utils"
)

type CommandType string

const (
	EXIT         CommandType = "exit"
	ECHO         CommandType = "echo"
	PWD          CommandType = "pwd"
	TYPE         CommandType = "type"
	USER_DEFINED CommandType = "user_defined"
)

func GetCommand(commandType CommandType) (Command, error) {
	switch commandType {
	case EXIT:
		return Exit{}, nil
	case ECHO:
		return Echo{}, nil
	case TYPE:
		return Type{}, nil
	case PWD:
		return Pwd{}, nil
	default:
		// check in path
		fileInPath, found := utils.FindFileInPath(string(commandType))
		if found {
			return UserDefined{
				Name: string(commandType),
				Path: fileInPath,
			}, nil
		}
		return nil, fmt.Errorf("%s: not found", string(commandType))
	}
}
