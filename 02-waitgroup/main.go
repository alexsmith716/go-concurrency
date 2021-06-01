package main

import (
	"fmt"
	"sync"
	"time"
)

// "sync" package provides 'WaitGroup' type for handling multiple goroutines
// 'WaitGroup' type can wait for a collection of goroutines to finish execution
// 'WaitGroup' is a 'counter' that tracks the total number of running goroutines

// 'WaitGroup' has 3 methods:
// 'Add()'  -increments 'WaitGroup' counter by specified 'arg' integer of new goroutines 
// 'Done()' -decrements 'WaitGroup' counter automatically per completed goroutine
// 'Wait()' -blocks the program until 'WaitGroup' counter reaches zero. final goroutine decrements to 0 & blocking removed

var urls = []string{
	"https://github.com",
	"https://gists.github.com",
	"https://www.googleapis.com",
}

// the 'WaitGroup' parameter points "*" to value at address "&wg"
func fetch(url string, wg *sync.WaitGroup) {
	fmt.Printf("Fetching url: %v \n", url)
	time.Sleep(3 * time.Second)
	fmt.Printf("Waited for response: %v \n", url)
	wg.Done()
}

func main() {

	// new WaitGroup instance
	var wg sync.WaitGroup

	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		// passing 'WaitGroup' memory address to 'fetch()'
		go fetch(urls[i], &wg)
	}

	wg.Wait()
}

//	% go run main.go
//	Fetching url: https://www.googleapis.com 
//	Fetching url: https://github.com 
//	Fetching url: https://gists.github.com 
//	Waited for response: https://gists.github.com 
//	Waited for response: https://github.com 
//	Waited for response: https://www.googleapis.com
