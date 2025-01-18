package main

import (
	"bufio"
	"fmt"
	"os"
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
		trimmedInput := commandInput[:len(commandInput)-1]
		if trimmedInput != "" {
			fmt.Fprintf(os.Stdout, "%s: command not found", trimmedInput)
		}
		fmt.Fprintln(os.Stdout)
	}
}
