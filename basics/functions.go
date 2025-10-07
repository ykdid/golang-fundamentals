package main

import (
	"fmt"
)

func FunctionsDemo() {

	greet()

	add(2,3)

	result := multiply(3, 4)
	fmt.Println("3 * 4 =", result)

	sum, diff := calculate(4,3)
	fmt.Println("4 + 3 =", sum, "4 - 3 =", diff)

	total := addWithNamedReturn(5, 6)
	fmt.Println("Total (named return):", total)

	fmt.Println("Sum of many:", sumAll(1, 2, 3, 4, 5))

	//function as a variable

	op := multiply
	fmt.Println("function as a variable: ", op(2,4))

	//anonym (lambda like) function

	func (message string)  {
		fmt.Println("Anonymous func says:", message)
	}("Hello!!!!!")

	counter := createCounter()
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())

}

//basic function
func greet() {
	fmt.Println("Hello!")
}

//with parameters no return
func add(a int, b int) {
	fmt.Println("Total: ", a + b)
}

// with return
func multiply(a int, b int) int {
	return a * b
}

//multiple returns

func calculate(a int, b int) (int, int){
	return a + b, a - b
}

//named return values
func addWithNamedReturn(a, b int) (sum int) {
	sum = a + b
	return // implicit return sum
}

//variadic function

func sumAll(numbers ...int) int{
	total := 0

	for _, num := range numbers{
		total += num
	}

	return total
}

//closure (function returns another function)
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
