package main

import (
	"fmt"
	"time"
)

// example demonstrates combining goroutines with channels and the resulting 'blocking action' of channels
// when goroutines send messages, the program may block progress while it waits to receive those messages
// a Go solution to that 'blocking action' is to 'listen' for 'receiving messages' from channels in a 'select' statement block
// a 'select' statement block chooses which of a set of possible 'send' or 'receive' 'channel operations' will proceed
// all select statement 'cases' refer to channel communication operations
// the 'select' statement blocks until one of its 'cases' finds a match, that match then gets executed

// REVIEW: >>>>> an intialized variable consists of 3 parts (name, value, memory address) <<<<<

// https://golang.org/ref/spec#Select_statements
// https://golang.org/pkg/time/#After

// "04:05" is a predefined 'Time.Format' layout for hour:minute

// send a 'fast' 1 sec delayed message
func fastChannelSender(channel chan string) {
	for {
		channel <- time.Now().Format("04:05")
		time.Sleep(1 * time.Second)
	}
}

// send a 'slow' 6 sec delayed message
func slowChannelSender(channel chan string) {
	for {
		channel <- time.Now().Format("04:05")
		time.Sleep(6 * time.Second)
	}
}


func main() {

	// create a 7 second counter channel
	timeoutChannel := time.After(26 * time.Second)

	// create 2 empty channels
	fastChannel := make(chan string)
	slowChannel := make(chan string)

	// goroutine function calls to the sending functions
	go fastChannelSender(fastChannel)
	go slowChannelSender(slowChannel)

	// a labeled break -a label to target the 'breaking-out' of the loop
	LabeledStatement:
	for {
		select {
		case done := <- timeoutChannel:
			// time expired, stop processing and transfer control/target the 'LabeledStatement'
			fmt.Println("================================================")
			fmt.Printf("time duration: '%T' >>>>> has elapsed at: %v \n", timeoutChannel, done.Format("04:05"))
			break LabeledStatement
		case slowChannelMessage := <- slowChannel:
			fmt.Printf("%v 	--slowChannelMessage Received! \n", slowChannelMessage)
		case fastChannelMessage := <- fastChannel:
			fmt.Printf("%v --fastChannelMessage Received! \n", fastChannelMessage)
		}
	}
}

// example with 'fastChannelSender()' sleeping 1 sec & 'slowChannelSender()' sleeping 1 sec
//
//	% go run main.go
//	16:22 	--slowChannelMessage Received! 
//	16:22 --fastChannelMessage Received! 
//	16:23 --fastChannelMessage Received! 
//	16:24 --fastChannelMessage Received! 
//	16:25 --fastChannelMessage Received! 
//	16:26 --fastChannelMessage Received! 
//	16:27 --fastChannelMessage Received! 
//	16:28 	--slowChannelMessage Received! 
//	16:28 --fastChannelMessage Received! 
//	16:29 --fastChannelMessage Received! 
//	16:30 --fastChannelMessage Received! 
//	16:31 --fastChannelMessage Received! 
//	16:32 --fastChannelMessage Received! 
//	16:33 --fastChannelMessage Received! 
//	16:34 	--slowChannelMessage Received! 
//	16:34 --fastChannelMessage Received! 
//	16:35 --fastChannelMessage Received! 
//	16:36 --fastChannelMessage Received! 
//	16:37 --fastChannelMessage Received! 
//	16:38 --fastChannelMessage Received! 
//	16:39 --fastChannelMessage Received! 
//	16:40 	--slowChannelMessage Received! 
//	16:40 --fastChannelMessage Received! 
//	16:41 --fastChannelMessage Received! 
//	16:42 --fastChannelMessage Received! 
//	16:43 --fastChannelMessage Received! 
//	16:44 --fastChannelMessage Received! 
//	16:45 --fastChannelMessage Received! 
//	16:46 	--slowChannelMessage Received! 
//	16:46 --fastChannelMessage Received! 
//	16:47 --fastChannelMessage Received! 
//	================================================
//	time duration: '<-chan time.Time' >>>>> has elapsed at: 16:48
