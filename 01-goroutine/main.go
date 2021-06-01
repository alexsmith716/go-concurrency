package main

import (
	"fmt"
	"time"
)

// 'goroutines' are concurrency in Go
// a goroutine is a lightweight thread of execution managed by the Go runtime
// Go executes code sequentially, line-by-line
// Go can also execute code concurrently with 'goroutines' (independent tasks executed asynchronously)
// Go concurrency takes advantage of multi-core processors
// every concurrently executing task in Go is a 'go' goroutine
// a goroutine is created placing the `go` keyword before a function or method call

// every Go program contains at least 1 goroutine, the "main goroutine / main function"
// the "main goroutine / main function" needs to be running for other goroutines to run
// by default, a Go program will exit when the "main goroutine" completes -regardless of any other goroutines
// so, the sequential flow of "control" of a Go program is synchronous with a "normal" function -waits for call to finish executing 
// with a "goroutine", "control" does not wait for call to finish executing, "control" immediately returns to next line of code 

func loopAndSleep(item string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%v %v   ", i, item)
		// 'Sleep()' will prevent the main program from terminating before the execution of the goroutines
		// 'Sleep()' will block the "main goroutine" until all other goroutines finish execution
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}

func main() {

	fmt.Println("\nLine-by-Line Execution...")
	loopAndSleep("Synch-Call-1")
	loopAndSleep("Synch-Call-2")

	fmt.Println("\nConcurrent Execution...")
	go loopAndSleep("Asynch-Call")
	loopAndSleep("Synch-Call-3")

	// time.Sleep(time.Second)
}

//	% go run main.go
//	
//	Line-by-Line Execution...
//	1 Synch-Call-1   2 Synch-Call-1   3 Synch-Call-1   
//	1 Synch-Call-2   2 Synch-Call-2   3 Synch-Call-2   
//	
//	Concurrent Execution...
//	1 Synch-Call-3   1 Asynch-Call   2 Synch-Call-3   2 Asynch-Call   3 Synch-Call-3   3 Asynch-Call
