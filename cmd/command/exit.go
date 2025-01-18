package command

import "os"

type Exit struct {
}

func (e Exit) Run(args []string) {
	exitCode := "0"
	if len(args) >= 1 {
		exitCode = args[0]
	}
	switch exitCode {
	case "0":
		os.Exit(0)
	default:
		os.Exit(1)
	}
}
