package main

import (
	"log"
	"sync"
	"time"
)

const (
	MaxJobs = 10
)

type Job struct {
	ID       int
	Number   int
	Attempts int
}

type Result struct {
	JobID    int
	Success  bool
	Result   int
	Error    error
	Attempts int
	Duration time.Duration
}

func main() {
	//Define number of tasks to do
	numTasks := 100

	//Create the tasks in an array
	var jobs []Job
	for v := 0; v < numTasks; v++ {
		job := Job{
			ID:     v,
			Number: v,
		}
		jobs = append(jobs, job)
	}

	//Create the channel with which to send the tasks to the workers
	jobsCh := make(chan Job, numTasks)

	//Create a channel to which the results are sent
	resultCh := make(chan Result)

	//Start sending the jobs by draining  the array created previously
	for _, job := range jobs {
		jobsCh <- job
	}
	close(jobsCh)

	//Spawn the maximum amount of worker so that they can start consuming the jobs
	var wg sync.WaitGroup
	for i := 0; i <= MaxJobs; i++ {
		wg.Add(1)
		go worker(jobsCh, &wg, resultCh)
	}
	wg.Wait()
}

func worker(jobsCh <-chan Job, wg *sync.WaitGroup, resultCh chan<- Result) {
	defer wg.Done()
	for job := range jobsCh {
		result := (job.Number * job.Number)
		log.Printf("Result: %v", result)
		solution := Result{
			JobID:    job.ID,
			Success:  true,
			Result:   result,
			Error:    nil,
			Attempts: 1,
			Duration: 1 * time.Second,
		}
		resultCh := solution
	}
}

func aggregator(resultCh <-chan Result) {

}
