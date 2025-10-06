package main

import "fmt"

func OperatorsDemo() {

	fmt.Println("Operators Demo:")

	a := 10
	b := 5

	var result int

	result = a + b
	fmt.Println("a + b =", result)

	result = a - b
	fmt.Println("a - b =", result)

	result = a * b
	fmt.Println("a * b =", result)

	result = a / b
	fmt.Println("a / b =", result)

	fmt.Println("Also use like this, a mod b = ", a % b)
}