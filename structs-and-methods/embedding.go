package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

//method on human
func (h Human) Greet() {
	fmt.Println("Hello,", h.Name)
}

//embedded struct -employee includes human

type Employee struct {
	Human //embedded
	Position string
	Department string
}

func (e Employee) Work() {
	fmt.Println(e.Name, "working as", e.Position, "in", e.Department)
}

func EmbeddingDemo() {

	fmt.Println("Embedding Demo:")

	emp := Employee {
		Human: Human{Name: "Yusuf", Age: 24},
		Position: "Software Developer",
		Department: "IT",
	}

	fmt.Println(emp.Name) //embedded field
	fmt.Println(emp.Position)
	fmt.Println(emp.Age) //embedded field

	emp.Work()
	emp.Greet() //embedded method
}
