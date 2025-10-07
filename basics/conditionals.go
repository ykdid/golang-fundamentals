package main

import "fmt"

func ConditionalsDemo() {

	fmt.Println("Conditionals Demo:")

	//basic if
	
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult.")
	}


	//if-else

	temp := 25

	if temp > 30 {
		fmt.Println("Greater than 30")
	} else{
		fmt.Println("Not Greater than 30")
	}

	//if - else if - else

	score := 85

	if score >=90 {
		fmt.Println("AWESOME")
	} else if score >= 80 {
		fmt.Println("GREAT")
	} else if score >= 70 {
		fmt.Println("NICE")
	} else {
		fmt.Printf("FAIL")
	}

	//declaration + condition

	if num := 5; num % 2 == 0 {
		fmt.Println("num is even")		
	} else {
		fmt.Println("num is odd")
	}

	//switch case

	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("1st day of week")
	case "Friday":
		fmt.Println("Almost weekend!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Midweek day")
	}

	// Switch without expression (acts like if-else chain)

	x := -5
	switch {
	case x > 0:
		fmt.Println("Positive number")
	case x < 0:
		fmt.Println("Negative number")
	default:
		fmt.Println("Zero")
	}
	
}