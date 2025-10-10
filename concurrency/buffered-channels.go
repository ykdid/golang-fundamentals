package main

import "fmt"

func BufferedChannelsDemo() {

	fmt.Println("Buffered Channels Demo:")

	ch := make(chan int, 2) // two capasity buffer

	ch <- 10
	ch <- 20

	fmt.Println("Sent two value to channel")

	//fifo
	fmt.Println(<-ch) //10
	fmt.Println(<-ch) //20
}