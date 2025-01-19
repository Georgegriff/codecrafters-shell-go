package command

import (
	"fmt"
	"os"
	"os/exec"
)

type UserDefined struct {
	Name string
	Path string
}

func (u UserDefined) CommandType() CommandType {
	return USER_DEFINED
}

func (u UserDefined) Run(args []string) {
	cmd := exec.Command(u.Name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("%s: command not found\n", u.Name)
	}
}

func (u UserDefined) String() string {
	return u.Name
}
