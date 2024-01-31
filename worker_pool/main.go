package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
The following are the core functionalities of our worker pool

Creation of a pool of Goroutines which listen on an input buffered channel waiting for jobs to be assigned
Addition of jobs to the input buffered channel
Writing results to an output buffered channel after job completion
Read and print results from the output buffered channel
*/

type Job struct {
	id        int
	randomNum int
}

type Result struct {
	job         Job
	sumOfDigits int
}

var (
	jobs    = make(chan Job, 10)
	results = make(chan Result, 10)
)

// finds the sum of individual digits of an integer
func calculateSumOfDigits(number int) int {
	sum := 0
	no := number

	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}

	// time.Sleep is used only to simulate the fact that calculation takes some time
	time.Sleep(2 * time.Second)

	return sum
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	close(results)
}

// creates a worker Goroutine
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{
			job:         job,
			sumOfDigits: calculateSumOfDigits(job.randomNum),
		}

		results <- output
	}
	wg.Done()
}

func allocateJobsToWorkers(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNum := rand.Intn(999)
		job := Job{
			id:        i,
			randomNum: randomNum,
		}
		jobs <- job
	}
	close(jobs)
}

// the result function reads the results channel and prints details
func result(done chan bool) {
	for res := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", res.job.id, res.job.randomNum, res.sumOfDigits)
	}

	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 100

	go allocateJobsToWorkers(noOfJobs)

	done := make(chan bool)
	go result(done)

	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
