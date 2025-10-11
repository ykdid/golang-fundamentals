package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const storageFile = "todos.json"

func SaveTodos (list TodoList) error {
	data, err := json.MarshalIndent(list, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(storageFile, data, 0644)
}

func LoadTodos () TodoList {
	var list TodoList
	data, err := os.ReadFile(storageFile)
	if err != nil {
		if os.IsNotExist(err) {
			return TodoList{}
		}
		fmt.Println("Error reading file:", err)
		return TodoList{}
	}

	json.Unmarshal(data, &list)
	return list
}