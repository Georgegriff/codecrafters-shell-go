package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	commandInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	trimmedInput := strings.TrimSpace(commandInput)
	if trimmedInput != "" {
		fmt.Fprintf(os.Stdout, "%s: command not found", trimmedInput)
	}
}
