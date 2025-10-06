package main

import "fmt"

func VariablesDemo() {
	
	fmt.Println("Variables Demo:")

	//basic declaration
	var name string = "Yusuf"
	var age int = 24
	fmt.Println("Name:", name, " Age:", age)

	//type inference
	var city = "Istanbul"
	fmt.Println(city)

	//short declaration
	country := "Turkiye"
	fmt.Println(country)

	//multiple declaration
	var x, y, z = 1, 2.5, "hello"
	fmt.Println(x, y, z)

	//zero values
	var number int
	var text string
	var isValid bool
	fmt.Println(number, text, isValid)

	//reassignment

	name = "kaya"
	fmt.Println(name)
}
