package command

import (
	"fmt"
	"os"
	"strings"
)

type Echo struct {
}

func (e Echo) Run(args []string) {
	fmt.Fprint(os.Stdout, strings.Join(args, " "))
}
