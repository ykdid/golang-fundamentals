package main

import (
	"fmt"
	"time"
)

func SelectDemo() {

	fmt.Println("Select Demo:")

	//created two channel unbuffered
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(){
		time.Sleep(200 * time.Millisecond)
		ch1 <- "Channel 1's message: I'm Channel 1."
	}()

	go func(){
		time.Sleep(100 * time.Millisecond)
		ch2 <- "Channel 2's message: I'm Channel 2."
	}()

	//select listens many channels , non-deterministic behavior, whichever channel sends data first, that case works.
	
	for i := 0; i<2; i++ {
		select{
		case msg1 := <-ch1:
			fmt.Println("Recieved:", msg1)
		
		case msg2 := <-ch2:
			fmt.Println("Recieved:", msg2)
		}
	}
}