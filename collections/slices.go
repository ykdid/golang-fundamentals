package main

import "fmt"


func SlicesDemo() {

	fmt.Println("Slices Demo:")

	//create a slice from array

	arr := [5]int {10, 20, 30, 40, 50}
	slice := arr[1:4] // includes index 1,2,3
	fmt.Println("Slice:", slice ,"Array:", arr)

	//create slice directly using literal

	nums := []int {1, 2, 3}
	fmt.Println("Numbers:", nums)

	//append elements

	nums = append(nums, 4, 5)
	fmt.Println("After append:", nums)

	more := []int {6, 7, 8}
	nums = append(nums, more...)
	fmt.Println("After append slice:", nums)

	// make slice with make()

	values := make([]string, 2, 5) //len=2 cap=5
	values[0] = "Hello"
	values[1] = "Go!"
	fmt.Println("Values:", values, "Len:", len(values), "Cap:", cap(values))

	//copy slices

	source := []int {1, 2, 3}
	dest := make([]int, len(source))
	copy(dest, source)
	fmt.Println("Coppied:", dest, "Original:", source)

	// nil vs empty slice

	var s1 []int //nil
	s2 := []int {}
	fmt.Println("s1 == nil ?", s1 == nil)
	fmt.Println("s2 == nil ?", s2 == nil)

	//slice iteration

	for i, v := range nums{
		fmt.Printf("Index %d -> Value %d\n", i, v)
	}
}