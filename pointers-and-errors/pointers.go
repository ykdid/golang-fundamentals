package main

import "fmt"

func PointersDemo() {

	fmt.Println("Pointers Demo:")

	//normal variable

	x := 10
	fmt.Println("x:", x)

	//pointer & -> address

	ptr := &x
	fmt.Println("Pointer:", ptr)
	fmt.Println("Pointer Value:", *ptr)

	//manipulating on pointer

	*ptr = 20
	fmt.Println("x after change:", x)

	//send ptr fn

	increment(&x)
	fmt.Println("x after change", x)
}

func increment(num *int){
	*num ++
}