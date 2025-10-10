package main

import(
	"fmt"

    "github.com/YusufKaya/golang-fundamentals/packages-and-modules/mathutil"
    "github.com/YusufKaya/golang-fundamentals/packages-and-modules/stringutil"
)

func main() {
		
	fmt.Println("Packages & Modules Demo:")

	sum := mathutil.Add(5, 7)
	fmt.Println("Add:", sum)

	product := mathutil.Multiply(3, 4)
	fmt.Println("Multiply:", product)

	word := "golang"
	fmt.Println("Reverse:", stringutil.Reverse(word))
	fmt.Println("ToUpper:", stringutil.ToUpper(word))
}
