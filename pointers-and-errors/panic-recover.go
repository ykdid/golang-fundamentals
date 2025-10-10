package main

import "fmt"

func PanicRecover() {

	fmt.Println("Panic Recover Demo:")

	defer func () {
		if r := recover(); r != nil { //recovers from panic generally using in defer
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Before panic")
	causePanic()
	fmt.Println("After panic") // won't work
}

func causePanic() {
	panic("something went wrong") // stops programme, uses in unexpected critical errors -in normal use error interface
}