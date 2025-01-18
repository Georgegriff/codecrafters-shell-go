package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		handleInput(trimmedInput)
		fmt.Fprintln(os.Stdout)
	}
}

func handleInput(command string) {
	fmt.Fprintf(os.Stdout, "%s: command not found", command)
}
