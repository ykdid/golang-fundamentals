package main

import "fmt"

func LoopsDemo() {

	fmt.Println("Loops Demo:")

	//basic for loop

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	//while-like loop

	count := 0

	for count < 3 {
		fmt.Println("Count is ", count)
		count++
	}

	//infinite loop

	num := 0

	for{
		if num >= 3 {
			break
		}

		fmt.Println("Infinite loop example: ", num)
		num ++
	}

	//continue for skip next iteration

	for i := 0; i <= 5; i++ {

		if i == 3 {
			fmt.Println("Skipping: ", i)
			continue
		}

		fmt.Println("i = ", i)
	}

	//range loops

	fruits := []string {"apple", "banana", "cherry"}

	for index, fruit := range fruits{
		fmt.Println("Index:", index, "Fruit:", fruit)
	}

	//if only need values

	for _, fruit := range fruits{
		fmt.Println("Fruit", fruit)
	}

	 //range with string (iterates over Unicode characters)
	for index, char := range "GoLang" {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}
}