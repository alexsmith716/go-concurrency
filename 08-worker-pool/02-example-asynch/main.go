package main

import (
	"fmt"
	"time"
	"sync"
)

// example demonstrates asychronously processing data (no worker pool)
// example demonstrates processing 100 simulated API calls each of which take 1/10th of a sec to complete
// this is in contrast to '02-example-synch' and '03-example-worker-pool'

type apiDataType struct {
	id int
}

func apiRequest(data apiDataType) {
	// fmt.Printf(">>>>>>>>> api %v request \n", data.id)
	time.Sleep(100 * time.Millisecond)
	// fmt.Printf("api %v response <<<<<<<<< \n", data.id)
}

func fetch(data apiDataType, wg *sync.WaitGroup) {
	defer wg.Done()
	apiRequest(data)
}

func work(allApiCalls []apiDataType) {
	fmt.Println("start simultaneously requesting 100 APIs ------------------")

	startTime := time.Now()

	// object is to load (pre-load) all API calls to measure elapsed time
	// since each API call is now its own goroutine, using 'WaitGroup' is needed
	// 'WaitGroup' will wait for all goroutines to finish (Promise.all())
	var wg sync.WaitGroup

	// create a buffered channel with a buffer size/capacity of 100
	// immediately write 100 integers into 'bufferedChannel' to match 'numApiCalls' (will avoid any blocking)
	// sends to a buffered channel are blocked only when the buffer is full
	bufferedChannel := make(chan apiDataType, 100)
	// bufferedChannel := make(chan apiDataType)
	// bufferedChannel := make(chan apiDataType, 3000)

	// add all 'allApiCalls' to 'bufferedChannel'
	for i := 0; i < len(allApiCalls); i++ {
		bufferedChannel <- allApiCalls[i]
	}

	// starting 'WaitGroup' collection
	// add 'main go func' goroutine to to 1st collection
	wg.Add(1)

	// self invoked goroutine will immediately read all data in the channel
	go func() {

		// close WaitGroup so it can be reused by following goroutines
		defer wg.Done()

		// while loop over channel data
		for {
			data, open := <- bufferedChannel
			if !open {
				break
			}

			// increment WaitGroup by each API call
			wg.Add(1)

			// each API call add is goroutine 
			go fetch(data, &wg)
		}
	}()

	close(bufferedChannel)

	wg.Wait()

	timeSinceStart := time.Since(startTime)

	fmt.Printf("total API processing time: %v \n", timeSinceStart)
}

func main() {

	// numApiCalls := 3000
	numApiCalls := 100

	// array of api data calls
	var allApiCalls []apiDataType

	// loop number of api calls and place into array
	for i := 0; i < numApiCalls; i++ {
		data := apiDataType{ id: i }
		allApiCalls = append(allApiCalls, data)
	}

	// call 'work' with all requests to process
	work(allApiCalls)
}

//	% go run mainB.go
//	start simultaneously requesting 100 APIs ------------------
//	total API processing time: 100.64489ms
