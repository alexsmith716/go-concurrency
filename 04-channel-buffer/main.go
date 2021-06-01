package main

import (
	"fmt"
)

// example demonstrates a 'buffered' channel (ability to send items to a channel as a 'queue')
// channels are 'unbuffered' by default (a receiver must exist for a sent message to be delivered)
// the 'sending action' is blocking, so sender & receiver must be in separate goroutines, otherwise `deadlock`
// specifying the 2nd arg 'buffer capacity' to the 'make()' function prevents 'unbuffered' `deadlock`

// https://golang.org/ref/spec#Channel_types
// https://golang.org/ref/spec#Making_slices_maps_and_channels

func main() {

	// create a buffered channel with a buffer size/capacity of 4
	// buffer size must be greater or equal to number of sent values (buffer overfill will cause 'deadlock')
	bufferedChannel := make(chan string, 4)

	// send a message to the channel
	bufferedChannel <- "Message one"

	// create a receiver for the sent channel massage
	bufferedChannelReceiver := <- bufferedChannel

	fmt.Printf("bufferedChannelReceiver: %v \n", bufferedChannelReceiver)

	// send 4 more messages to the 'bufferedChannel'
	bufferedChannel <- "Message two"
	bufferedChannel <- "Message three"
	bufferedChannel <- "Message four"
	bufferedChannel <- "Message five"

	// close the channel after all messages have been sent
	close(bufferedChannel)

	fmt.Println("len(bufferedChannel): ", len(bufferedChannel))

	for {
		bufferedChannelReceiver, open := <- bufferedChannel
		if !open {
			break
		}
		fmt.Println(bufferedChannelReceiver)
	}

	//	for bufferedChannelReceiver := range bufferedChannel {
	//		fmt.Println(bufferedChannelReceiver)
	//	}
}

//	% go run main.go
//	bufferedChannelReceiver: Message one 
//	len(bufferedChannel):  4
//	Message two
//	Message three
//	Message four
//	Message five
