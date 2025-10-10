package main

import(
	"fmt"
	"errors"
) 

func divide (a,b float64) (float64, error){
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func ErrorHandlingDemo () {

	fmt.Println("\nError Handling Demo:")

	//valid case

	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Result:", result)
	}

	//error case

	result, err = divide(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}


}