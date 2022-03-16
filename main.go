package main

import (
	"database/sql"
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
	db, err := CreateConnection()
	if err != nil {
		panic(err)
	}
	db.Exec("create table if not exists JOBCOMPANY(id integer, company char, age integer, address char, salary integer)")
	// todo: print the subcommand supplied by the user: X
	// todo: if no sub command provided, warn user: R
	// todo: if an unknown command is provided, warn user: R
	// todo: make a func for each subcommand of the same name like func create() etc and call them appropriately: R
	// todo: define every subcommand func to print "<command-name> not implemented": R
	// todo: create func for db calls

	if len(os.Args) > 1 {
		subCommand := os.Args[1]

		command := []string{"create", "get", "find", "update", "delete"}

		if !CommandHas(command, subCommand) {
			fmt.Println("Please provide a valid command. Available commands are create,get,find,update,delete")
		} else {
			switch subCommand {
			case "create":
				Create(db)
			case "get":
				Get(db)
			case "find":
				Find(db)
			case "update":
				Update(db)
			case "delete":
				Delete(db)
			}
		}
	} else {
		fmt.Println("Please provide a Subcommand")
	}
}

func Create(db *sql.DB) {
	// user input
	// input sanitize
	// db send
	// CreatePost(db, jp)
}

func Get(db *sql.DB) {
	// id := os.args[2]
	// jp := GetPost(db, id)
}
func Find(db *sql.DB) {
	// FindPost(db, company)
}
func Update(db *sql.DB) {
	// 	previousjp := GetPost(id)
	// 	UpdatePost(db, id, jp)
}
func Delete(db *sql.DB) {
	// DeletePost(db, id)
}
