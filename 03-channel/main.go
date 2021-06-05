package main

import (
	"fmt"
	"time"
)

// 'channels' enable goroutines to communicate with each other
// 'channels' send and receive messages -a blocking action
// sending a message:   the channel waits until the receiver can accept the message
// receiving a message: the channel waits until the sender delivers the message 
// channel operation is synchronous when sending and receiving messages

// channel is created with 'chan' keyword and a data type as arg to 'make()' function

// channelName := make(chan dataType)

// send a value to the channel:    "channelName <- message"
// receive a value from a channel: "variableName <- channelName"

// 'close()' function ensures a receiver is not left waiting for a message
// 'close()' function takes a channel name to close when channel has completed sending

// receiver receives a second boolean value indicating open state of channel (open (true) / closed (false))

// function creates a channel sender for 'newChannel' which sends 3 messages at a 1-sec interval and then closes
func sendTimeMessage(msg string, channel chan string) {

	// fmt.Printf("channel data type: %T \n", channel)

	for i := 0; i < 3; i++ {
		// 'send' a message with the time to channel 'newChannel'
		channel <- msg + time.Now().Format("04:05")
		time.Sleep(1 * time.Second)
	}
	close(channel)
	fmt.Println("Channel Closed --------------------------------")
}


func main() {

	// create empty channel
	newChannel := make(chan string)

	// add a goroutine that passes a channel
	go sendTimeMessage("Sending time message: ", newChannel)

	// loop over the 'open' channel 'receiver' and display received messages
	for {
		// break loop if channel state closes
		msg, open := <- newChannel
		if !open {
			break
		}
		fmt.Printf("%v --Message Received! \n", msg)
	}
}

//	% go run main.go
//	Sending time message: 43:54 --Message Received! 
//	Sending time message: 43:55 --Message Received! 
//	Sending time message: 43:56 --Message Received! 
//	Channel Closed --------------------------------
