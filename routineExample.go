package main

import (
	"fmt"
	"time"
)

type myChannels struct {
	stringChannel chan string
	intChannel chan int
	exitChannel chan bool
}

func waitingRoutine(channels myChannels) {
	for {
		fmt.Println("Waiting for a message")
		time.Sleep(2 * time.Second)
		fmt.Println("Reading a message")
		select {
		case number := <-channels.intChannel:
			println("Received this number:", number)
		case str := <-channels.stringChannel:
			println("Received this string:", str)
		case <-channels.exitChannel:
			fmt.Println("Closing the routine")
			return
		}
	}
}

func launchRoutine() {
	myChannels := myChannels{make(chan string), make(chan int), make(chan bool)}
	go waitingRoutine(myChannels)
	fmt.Println("Sending a number message")
	myChannels.intChannel <- 20
	fmt.Println("Sending a string message")
	myChannels.stringChannel <- "Hello world"
	fmt.Println("Sending the quit bool")
	myChannels.exitChannel <- true
	fmt.Println("Exiting the program")

}