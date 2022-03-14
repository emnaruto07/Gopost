package main

import (
	"fmt"
	"os"
)

func SliceHas(hay []string, needle string) bool {
	// hay = [a, d, f, g], needle = f, true

	for _, i := range hay {
		if i == needle { //
			return true
		}
	}
	return false
}

func main() {
	// todo: print the subcommand supplied by the user: X
	// todo: if no sub command provided, warn user
	// todo: if an unknown command is provided, warn user

	if len(os.Args) > 1 {
		// subCommand := os.Args[1]

		// var command [5]string
		command := []string{"create", "get", "find", "update", "delete"}
		SliceHas(command, "get")

	} else {

		fmt.Println("Please provide a command")
	}
}
