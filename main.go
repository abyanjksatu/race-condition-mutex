package main

import (
	"fmt"
	"sync"
)

// Define global variable number
var number int

// Define mutex to safely access data across multiple goroutines
var mutex = &sync.Mutex{}

func main() {

	// Use goroutine for a lightweight thread of execution
	// with running Asynchronously to execute addMe() func
	go func() {
		addMe(1)
	}()

	// Use goroutine for a lightweight thread of execution
	// with running Asynchronously to execute addMe() func
	go func() {
		addMe(2)
	}()
}

// This function is to add number sequentially
// to avoid race condition with mutex
func addMe(id int) {

	// Lock() the mutex to ensure exclusive access
	// to the state, read the value at the chosen key
	mutex.Lock()

	// Unlock() he mutex, and increment the readOps count
	defer mutex.Unlock()

	// Add the id to number
	number = number + id

	// print the last number was added
	fmt.Println(number)
}
