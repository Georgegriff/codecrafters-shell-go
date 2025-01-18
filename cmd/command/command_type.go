package command

import (
	"fmt"
)

type CommandType string

const (
	EXIT CommandType = "exit"
	ECHO CommandType = "echo"
	TYPE CommandType = "type"
)

func GetCommand(commandType CommandType) (Command, error) {
	switch commandType {
	case EXIT:
		return Exit{}, nil
	case ECHO:
		return Echo{}, nil
	case TYPE:
		return Type{}, nil
	default:
		return nil, fmt.Errorf("%s: not found", string(commandType))
	}
}
