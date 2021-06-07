package main

import (
	"fmt"
	"time"
)

// example demonstrates sychronously processing data
// example demonstrates processing 100 simulated API calls each of which take 1/10th of a sec to complete
// this is in contrast to '02-example-asynch' and '03-example-worker-pool'


type apiDataType struct {
	id int
}

func work(allApiCalls []apiDataType) {
	fmt.Println("start simultaneously requesting 100 APIs ------------------")

	startTime := time.Now()

	// do not need the element value "_"
	//	for i, _ := range allApiCalls {
	//		apiRequest(allApiCalls[i])
	//	}

	for i := 0; i < len(allApiCalls); i++ {
		apiRequest(allApiCalls[i])
	}

	timeSinceStart := time.Since(startTime)

	// display total processing time 
	fmt.Printf("total API processing time: %v \n", timeSinceStart)
}

// api request delay of 1 sec
func apiRequest(data apiDataType) {
	// fmt.Printf(">>>>>>>>> api %v request \n", data.id)
	time.Sleep(100 * time.Millisecond)
	// fmt.Printf("api %v response <<<<<<<<< \n", data.id)
}

func main() {

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

//	% go run main.go
//	start simultaneously requesting 100 APIs ------------------
//	total API processing time: 10.364880484s
