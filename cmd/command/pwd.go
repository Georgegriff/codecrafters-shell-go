package command

import (
	"fmt"
	"os"
)

type Pwd struct {
}

func (c Pwd) CommandType() CommandType {
	return PWD
}

func (c Pwd) String() string {
	return string(c.CommandType())
}

func (c Pwd) Run(args []string) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprint(os.Stdout, err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "%s\n", wd)
}
