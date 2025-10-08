package main

import "fmt"

func ArraysDemo() {

	fmt.Println("Arrays Demo:")

	// basic array

	var numbers [3]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	fmt.Println("Numbers:", numbers)
	 
	//array literal

	colors := [3]string {"red", "white", "black"}
	fmt.Println("Colors:", colors)

	primes := [...]int{2, 3, 5, 7, 11}
	fmt.Println("Primes:", primes, "Length:", len(primes))

	//loop throug array

	for index, value := range colors {
		fmt.Println("Index:", index, "Value:", value)
	}

	//multi dimensional array

	matrix := [2][3]int{
		{1,2,3},
		{4,5,6},
	}

	fmt.Println("Matrix:", matrix)
	fmt.Println("Access element [1][2]", matrix[1][2])
}