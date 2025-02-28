package command

import (
	"fmt"

	"github.com/codecrafters-io/shell-starter-go/cmd/utils"
)

type CommandType string

const (
	CAT          CommandType = "cat"
	CD           CommandType = "cd"
	EXIT         CommandType = "exit"
	ECHO         CommandType = "echo"
	PWD          CommandType = "pwd"
	TYPE         CommandType = "type"
	USER_DEFINED CommandType = "user_defined"
)

func GetCommand(commandType CommandType) (Command, error) {
	switch commandType {
	// Disable cat for now because codecrafters test fail with it
	// case CAT:
	// 	return Cat{}, nil
	case EXIT:
		return Exit{}, nil
	case ECHO:
		return Echo{}, nil
	case TYPE:
		return Type{}, nil
	case PWD:
		return Pwd{}, nil
	case CD:
		return Cd{}, nil
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
