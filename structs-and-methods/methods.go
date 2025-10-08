package main

import "fmt"


// struct definition

type Rectangle struct{
	Width float64
	Height float64
}

//value receiver method

func (r Rectangle) Area() float64 {

	return r.Width * r.Height
}


func (r *Rectangle) Scale(factor float64) {

	r.Width = r.Width * factor
	r.Height = r.Height *factor
}

type Person struct {
	Name string
	Age int
}

func (p Person) Greet() {
	fmt.Println("Hello Go! You are not difficult, ez!!")
}

func MethodsDemo () {

	fmt.Println("Methods Demo:")

	rect := Rectangle{Width: 50, Height: 5}
	fmt.Println("Area:", rect.Area())

	rect.Scale(2)
	fmt.Println("Scaled Rectangle:", rect)
	fmt.Println("New Area:", rect.Area())

	p1 := Person{"Yusuf", 24}
	p1.Greet()
}

