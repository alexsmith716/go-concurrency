package main

import (
	"fmt"
	"time"
)
 
// example demonstrates combining goroutines with channels and the resulting 'blocking action' of channels
// when goroutines send messages, the program may block progress while it waits to receive those messages (the slower ones)
// this example is in contrast to the 'select statement' example

// https://golang.org/ref/spec#Select_statements

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

	// create 2 channels
	fastChannel := make(chan string)
	slowChannel := make(chan string)

	// goroutine function calls to the sending functions
	go fastChannelSender(fastChannel)
	go slowChannelSender(slowChannel)

	// loop with receiver for each channels
	for {
		slowChannelMessage, open := <- slowChannel
		fastChannelMessage, open := <- fastChannel
		if !open {
			break
		}
		fmt.Println(">>>>>>>>>> Received Messages at:", time.Now().Format("04:05"), "<<<<<<<<<<<<<")
		fmt.Printf("%v --fastChannelMessage Received! \n", fastChannelMessage)
		fmt.Printf("%v 	--slowChannelMessage Received! \n", slowChannelMessage)
	}

	// block forever
	select {}
}

// example with 'fastChannelSender()' sleeping 1 sec & 'slowChannelSender()' sleeping 6 secs
//
//	% go run main.go
//	>>>>>>>>>> Received Messages at: 15:04 <<<<<<<<<<<<<
//	15:04 --fastChannelMessage Received! 
//	15:04 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 15:10 <<<<<<<<<<<<<
//	15:05 --fastChannelMessage Received! 
//	15:10 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 15:16 <<<<<<<<<<<<<
//	15:11 --fastChannelMessage Received! 
//	15:16 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 15:22 <<<<<<<<<<<<<
//	15:17 --fastChannelMessage Received! 
//	15:22 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 15:28 <<<<<<<<<<<<<
//	15:23 --fastChannelMessage Received! 
//	15:28 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 15:34 <<<<<<<<<<<<<
//	15:29 --fastChannelMessage Received! 
//	15:34 	--slowChannelMessage Received! 
//	^Csignal: interrupt

// example with 'fastChannelSender()' sleeping 1 sec & 'slowChannelSender()' sleeping 1 sec
//
//	% go run main.go
//	>>>>>>>>>> Received Messages at: 17:29 <<<<<<<<<<<<<
//	17:29 --fastChannelMessage Received! 
//	17:29 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 17:30 <<<<<<<<<<<<<
//	17:30 --fastChannelMessage Received! 
//	17:30 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 17:31 <<<<<<<<<<<<<
//	17:31 --fastChannelMessage Received! 
//	17:31 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 17:32 <<<<<<<<<<<<<
//	17:32 --fastChannelMessage Received! 
//	17:32 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 17:33 <<<<<<<<<<<<<
//	17:33 --fastChannelMessage Received! 
//	17:33 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 17:34 <<<<<<<<<<<<<
//	17:34 --fastChannelMessage Received! 
//	17:34 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 17:35 <<<<<<<<<<<<<
//	17:35 --fastChannelMessage Received! 
//	17:35 	--slowChannelMessage Received! 
//	^Csignal: interrupt

// example with 'fastChannelSender()' sleeping 6 sec & 'slowChannelSender()' sleeping 1 sec
//
//	% go run main.go
//	>>>>>>>>>> Received Messages at: 18:49 <<<<<<<<<<<<<
//	18:49 --fastChannelMessage Received! 
//	18:49 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 18:55 <<<<<<<<<<<<<
//	18:55 --fastChannelMessage Received! 
//	18:50 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 19:01 <<<<<<<<<<<<<
//	19:01 --fastChannelMessage Received! 
//	18:51 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 19:07 <<<<<<<<<<<<<
//	19:07 --fastChannelMessage Received! 
//	18:56 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 19:13 <<<<<<<<<<<<<
//	19:13 --fastChannelMessage Received! 
//	19:02 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 19:19 <<<<<<<<<<<<<
//	19:19 --fastChannelMessage Received! 
//	19:08 	--slowChannelMessage Received! 
//	>>>>>>>>>> Received Messages at: 19:25 <<<<<<<<<<<<<
//	19:25 --fastChannelMessage Received! 
//	19:14 	--slowChannelMessage Received! 
//	^Csignal: interrupt
