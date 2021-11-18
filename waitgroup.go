package main

import (
	"fmt"
	"sync"
	"time"
)

// Create Wait Group
var wg sync.WaitGroup

func main() {

	// Add two jobs to the Wait Group
	wg.Add(2)

	// Launch both workers concurrently
	go worker1()
	go worker2()

	// Wait for all jobs in the Wait Group to complete
	wg.Wait()

	fmt.Println("All workers Done!")
}

func worker1() {
	// Decrement the Wait Group count, signaling a job is complete
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Worker 1")
		time.Sleep(1000 * time.Millisecond)
	}
}

func worker2() {
	// Decrement the Wait Group count, signaling a job is complete
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Worker 2")
		time.Sleep(2000 * time.Millisecond)
	}
}
