package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/command"
)

func main() {

	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		commandInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		trimmedInput := strings.TrimSuffix(commandInput, "\n")
		command.ExecuteCommandInput(trimmedInput)

		fmt.Fprintln(os.Stdout)
	}
}
