package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
	fmt.Println(`Usage:
	go run . <command> [arguments]

	Available commands:
	add <task>       Add a new task to your TODO list
	list             Show all tasks with their status
	done <task_id>   Mark a task as completed

	Examples:
	go run . add "Finish Go CLI project"
	go run . list
	go run . done 1`)
		return
	}

	command := os.Args[1]
	list := LoadTodos()

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a todo title.")
			return
		}

		title := os.Args[2]
		list.Add(title)
		SaveTodos(list)

	case"list":
		list.List()

	case"done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the todo ID.")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID.")
			return
		}
		list.Complete(id)
		SaveTodos(list)

	default:
		fmt.Println("Unknown command. Use add, list, or done.")
	}

}


