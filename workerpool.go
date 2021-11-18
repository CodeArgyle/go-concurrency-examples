package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Create channels
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Initialize random number generator (Used to sleep workers)
	rand.Seed(time.Now().UnixNano())

	// Create 3 workers
	for WId := 1; WId <= 3; WId++ {
		go work(WId, jobs, results)
	}

	// Send 10 'jobs' to the jobs channel
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	// Tell functions reading off the 'jobs' channel that there will be no more jobs
	close(jobs)

	// Read results from the results channel
	for result := range results {
		fmt.Printf("Job %d Complete!\n", result)
	}
	fmt.Println("All Jobs Complete!")

}

//work keeps falling a sleep on the job...
func work(WId int, jobs <-chan int, results chan<- int) {

	// Close the results channel when there are no more results to send
	defer close(results)

	// Read jobs off of the jobs channel
	for job := range jobs {
		// 'Work' for a random amount of time < 3s multiplied by the worker ID
		workTime := rand.Intn(3000) * WId
		fmt.Printf("Worker %d working on job %d for %d Milliseconds...\n", WId, job, workTime)
		// 'Work'
		time.Sleep(time.Duration(workTime) * time.Millisecond)
		// Send job to the results channel to signal job complete
		results <- job
	}
}
