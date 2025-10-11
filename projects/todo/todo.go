package main

import "fmt"

type Todo struct{
	ID int
	Title string
	Completed bool
}

type TodoList struct {
	Items []Todo
}

func (t *TodoList) Add (title string) {
	id := len(t.Items) + 1 
	newTodo := Todo{ID: id, Title: title, Completed: false}
	t.Items = append(t.Items, newTodo)
	fmt.Printf("Added: \"%s\"\n", title)
}

func (t *TodoList) List() {
	if len(t.Items) == 0 {
		fmt.Println("No todos found.")
		return
	}
	for _, todo := range t.Items {
		status := " "
		if todo.Completed {
			status = "✔"
		}
		fmt.Printf("[%s] %d: %s\n", status, todo.ID, todo.Title)
	}
}

func (t *TodoList) Complete(id int) {
	for i, todo := range t.Items{
		if todo.ID == id {
			t.Items[i].Completed = true
			fmt.Printf("Completed: \"%s\"\n", todo.Title)
			return
		}
	}

	fmt.Println("Todo not found.")
}