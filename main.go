package main

import (
	"fmt"
	"os"
)

func CommandHas(hay []string, needle string) bool {
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
	// todo: if no sub command provided, warn user: R
	// todo: if an unknown command is provided, warn user: R
	// todo: make a func for each subcommand of the same name like func create() etc and call them appropriately: R
	// todo: define every subcommand func to print "<command-name> not implemented"

	if len(os.Args) > 1 {
		subCommand := os.Args[1]

		// var command [5]string
		command := []string{"create", "get", "find", "update", "delete"}

		if !CommandHas(command, subCommand) { // command invalid
			fmt.Println("Please provide a valid command. Available commands are create,get,find,update,delete")
		} else {
			// if subCommand == "create"{
			// 	Create()
			// }
			switch subCommand {
			case "create":
				Create()
			case "get":
				Get()
			case "find":
				Find()
			case "update":
				Update()
			case "delete":
				Delete()
			}
		}
	} else {
		fmt.Println("Please provide a Subcommand")
	}
}

func Create() {
	// user input
	// input sanitize
	// db send
	fmt.Println("Create not implemented")
}
func Get() {
	fmt.Println("Get not implemented")
}
func Find() {
	fmt.Println("Find not implemented")
}
func Update() {
	fmt.Println("Update not implemented")
}
func Delete() {
	fmt.Println("Delete not implemented")
}
