package main

import "fmt"

func OperatorsDemo() {

	fmt.Println("Operators Demo:")

	a := 10
	b := 5
	
	//arithmetic
	var result int

	result = a + b
	fmt.Println("a + b =", result)

	result = a - b
	fmt.Println("a - b =", result)

	result = a * b
	fmt.Println("a * b =", result)

	result = a / b
	fmt.Println("a / b =", result)

	result = a % b
	fmt.Println("a mod b = ", result)

	//comparison
	fmt.Println("a == b ?", a == b)
	fmt.Println("a != b ?", a != b)
	fmt.Println("a > b ?", a > b)
	fmt.Println("a < b ?", a < b)
	fmt.Println("a >= b ?", a >= b)
	fmt.Println("a <= b ?", a <= b)

	//logical
	x := true
	y := false
	fmt.Println("x && y:", x && y)
	fmt.Println("x || y:", x || y)
	fmt.Println("!x: ", !x)

	//assignment
	n := 5
	n += 3
	fmt.Println("n += 3 -> " , n)
	n *= 2
	fmt.Println("n *=2 -> ", n)
}