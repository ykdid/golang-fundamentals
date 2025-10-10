package main

import "fmt"

func ChannelsDemo() {

	fmt.Println("Channels Demo:")

	ch := make(chan string) //unbuffered channel

	go func() {
		ch <- "Message from channel" //send data
	}()

	msg := <-ch
	fmt.Println(msg)
}

