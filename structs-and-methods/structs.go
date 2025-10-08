package main

import "fmt"

type User struct{
	Id int
	Name string
	Age int
	City string
}

type Product struct{
	Id int
	Title string
	Description string
	Price int
}

func StructsDemo() {

	fmt.Println("Structs Demo:")

	//basic initialization
	var u1 User
	u1.Id = 1
	u1.Name = "Yusuf"
	u1.Age = 24
	u1.City = "Wien"
	fmt.Println(u1)

	//struct literal

	u2 := User{Id: 2, Name: "Ali", Age: 29, City: "Istanbul"}
	fmt.Println(u2)

	//positional init -not recommended if struct changes often-

	u3 := User {3, "Ayşe", 24, "Yozgat"}
	fmt.Println(u3)

	//accessing

	fmt.Println("User", u1.Id,": Name", u1.Name)
	fmt.Println("Age", u3.Age)

	// pointer to struct

	u4 := &User{Id:4, Name: "Elif", Age: 22, City: "Bursa"}
	fmt.Println(u4)
	fmt.Println("Pointer access:", u4.Name) // Go automatically dereferences

	// anonymous struct

	student := struct {
		Name  string
		Score int
	}{
		Name:  "Merve",
		Score: 95,
	}
	fmt.Println(student)

}