package main

import (
	"fmt"
	"os"
	"strconv"
)

func main () {
if len(os.Args) < 2 {
	fmt.Println(`Usage:
	go run . [operation | num1 | num2]
	operation types: [add, sub, mul, div]
	examples: add 2 3, sub 4 1`)
		return
	}

	command := os.Args[1]

	num1, err:= strconv.Atoi(os.Args[2])
	if err != nil {
        panic(err)
    }

	num2, err:= strconv.Atoi(os.Args[3])
	if err != nil {
        panic(err)
    }

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please enter valid input.")
			return
		}
		fmt.Println(Add(num1, num2))

	case"sub":
		if len(os.Args) < 3 {
			fmt.Println("Please enter valid input.")
			return
		}
		fmt.Println(Sub(num1, num2))

	case"mul":
		if len(os.Args) < 3 {
			fmt.Println("Please enter valid input.")
			return
		}
		
		fmt.Println(Multiply(num1, num2))

	case"div":
		if len(os.Args) < 3 {
			fmt.Println("Please enter valid input.")
			return
		}
		fmt.Println(Divide(num1, num2))

	default:
		fmt.Println("Unknown command. Use add, sub, mul or div.")
	}

}