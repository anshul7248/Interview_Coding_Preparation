// Implement a worker pool using goroutines and channels.
package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	WorkerId int
	JobID    int
	Output   int
}

func main() {
	numJobs := 10
	numWorkers := 3
	jobs := make(chan int, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// Start Workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send Jobs

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs) //if i will not close this channel worker will blocks for forever

	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		fmt.Printf("Job %d prcoess by Worker %d -- Result  = %d\n",
			result.JobID,
			result.WorkerId,
			result.Output,
		)
	}
}

func worker(id int, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished %d job\n", id, job)

		results <- Result{
			WorkerId: id,
			JobID:    job,
			Output:   job * 2,
		}
	}
}
