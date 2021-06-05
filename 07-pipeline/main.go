package main

import (
	"fmt"
	"strconv"
)

// example demonstrates using the blocking (synchronizing) action of unbuffered channels to create a connected 'pipeline' of goroutines
// unbuffered channels are used to directly connect the goroutines together (the value sent from 'goroutineA' is the value received by 'goroutineB')
// the unbuffered channel causes each goroutine in the 'pipeline' to wait until a message is fully sent or received end-to-end

// channels can both send & receive and can also be 'directed' to ONLY send or receive

// SEND ONLY CHANNEL:    'func functionName (channelName <-chan channelDataType)'  <<<<<<<<<  "<-chan"
// RECEIVE ONLY CHANNEL: 'func functionName (channelName chan<- channelDataType)'  <<<<<<<<<  "chan<-"

// >>>>>> specifying a channel direction is not required but if specified, must be used correctly <<<<<<<

// https://blog.golang.org/pipelines

// >>>>>> DONT FORGET TO CLOSE FINISHED SENDING CHANNEL <<<<<<<

// channel 'numbers' may only 'receive' items
func startPipelineFunction(numbers chan<- int) {
	for i := 1; i <= 10; i++ {
		numbers <- i
	}
	close(numbers)
}

// channel 'numbers' is sent-only and 'squared' is receive-only
func continuePipelineFunctionA(numbers <-chan int, squared chan<- string) {
	for {
		res, open := <- numbers
		if !open {
			break
		}
		squared <- strconv.Itoa(res) + " is " + strconv.Itoa(res * res) // send 'numbers' to receiving 'squared' when pipeline is un-sync'd
	}
	close(squared)
}

func continuePipelineFunctionB(squared <-chan string, result chan<- string) {
	for {
		res, open := <- squared
		if !open {
			break
		}
		result <- "The square root of " + res  // send 'squared' to receiving 'result' when pipeline is un-sync'd
	}
	close(result)
}


func main() {

	// 3 empty, unbuffered channels
	numbers := make(chan int)
	squared := make(chan string)
	result := make(chan string)

	// start sequence of calling goroutines
	go startPipelineFunction(numbers) // begin pipeline with 'received' data
	go continuePipelineFunctionA(numbers, squared) // continue pipeline with previous channel 'sending' into new channel
	go continuePipelineFunctionB(squared, result) // continue pipeline with previous channel 'sending' into new channel

	// loop and disply ending channel pipeline data
	for {
		response, open := <- result
		if !open {
			break
		}
		fmt.Printf("%v \n", response)
	}
}

//	% go run main.go
//	The square root of 1 is 1 
//	The square root of 2 is 4 
//	The square root of 3 is 9 
//	The square root of 4 is 16 
//	The square root of 5 is 25 
//	The square root of 6 is 36 
//	The square root of 7 is 49 
//	The square root of 8 is 64 
//	The square root of 9 is 81 
//	The square root of 10 is 100
