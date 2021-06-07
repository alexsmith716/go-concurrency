package main

import (
	"fmt"
	"time"
	"sync"
)

// example demonstrates asychronously processing data employing a 'Worker Pool/Thread Pool'
// example demonstrates processing 1000 simulated API calls each of which take 1/10th of a sec to complete
// this is in contrast to '01-example-synch' and '02-example-asynch'

// a 'worker-pool' is a collection of threads that are waiting for assigned tasks
// a set 'pool' of workers processes tasks one after another which distributes tasks over time
// tasks are NOT processed all at once which could overload computer memory & CPU 
// for huge amount of data processing a computer can run out of memory (running several million tasks)
// the benefit of a 'worker-pool' is it distributes the process load over time (enables the handling of much greater workloads)

type apiDataType struct {
	id int
}

func apiRequest(data apiDataType) {
	// fmt.Printf(">>>>>>>>> api %v request \n", data.id)
	time.Sleep(100 * time.Millisecond)
	// fmt.Printf("api %v response <<<<<<<<< \n", data.id)
}

func workerPool(allApiCalls []apiDataType, numberOfWorkers int) {
	fmt.Println("start simultaneously requesting 100 APIs ------------------")

	startTime := time.Now()

	var wg sync.WaitGroup

	// use the buffered channel to implement the worker pool
	// a pool of goroutines ('numberOfWorkers') will listen on the buffered channel for assigned tasks
	bufferedChannel := make(chan apiDataType, numberOfWorkers)

	// for-loop all 'numberOfWorkers' >>> goroutines <<< to do processing (100 goroutines will process all 'allApiCalls')
	// once a worker/goroutine is done processig a task, it processes another
	// the channel is a queue and each worker (goroutine) is continually used to process as many assigned tasks as possible
	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// while loop all open 'bufferedChannels' and task each with 'apiRequest(data)'
			for {
				data, open := <- bufferedChannel
				if !open {
					break
				}
				apiRequest(data)
			}
		}()
	}

	// goroutines have now read/extracted (some) data in the channel (avoiding deadlock/panic)
	// data has been read/extracted from the channel
	// so now writing number of read/extracted 'allApiCalls' to 'bufferedChannel'
	// this read/write cycle continues until the channel is closed (data all processed)
	for i := 0; i < len(allApiCalls); i++ {
		bufferedChannel <- allApiCalls[i]
	}

	close(bufferedChannel)

	wg.Wait()

	timeSinceStart := time.Since(startTime)

	fmt.Printf("total API processing time: %v \n", timeSinceStart)
}

func main() {

	numApiCalls := 1000
	numberOfWorkers := 100

	var allApiCalls []apiDataType

	for i := 0; i < numApiCalls; i++ {
		data := apiDataType{ id: i }
		allApiCalls = append(allApiCalls, data)
	}

	workerPool(allApiCalls, numberOfWorkers)
}

//	% go run main.go
//	start simultaneously requesting 100 APIs ------------------
//	total API processing time: 1.024139554s
