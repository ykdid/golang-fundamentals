package main

import (
	"fmt"
	"time"
)

func sayHello() {
	
	for i := 1; i<=3; i++{
		fmt.Println("Hello from goroutine:", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func GoroutineDemo() {

	fmt.Println("Goroutine Demo:")

	go sayHello() //starts goroutine (as like thread but lighter) simultaneously

	for i := 1; i<=3; i++ {
		fmt.Println("Hello from main:", i)
		time.Sleep(250 * time.Millisecond)
	}

	time.Sleep(1 * time.Microsecond)
}
