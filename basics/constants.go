package main

import "fmt"

func ConstantsDemo() {

	fmt.Println("Constants Demo:")

	const Pi = 3.14159265458979323846264338327950288
	const AppName = "Yusuf's Go Journay"

	fmt.Println(Pi, AppName)

	const (
		Monday = 1
		Sunday = 2
		Tuesday = 3
	)

	fmt.Println(Monday, Sunday, Tuesday)

	//iota auto iterate like enum 0,1,2,..,n
	const (
		Red = iota
		White
		Black
	)

	fmt.Println("Colors:",Red, White, Black)
}