package main

import (
	"flag"
	"fmt"
	// "math/rand"
	"runtime"
	"sync"
	"time"
)

const defaultPendingJobsNo = 50

var (
	// t map[int]int
	limit  chan interface{}
	jobsNo *int
	wg     sync.WaitGroup
)

func init() {
	limit = make(chan interface{}, runtime.NumCPU())
	jobsNo = flag.Int("jobs", defaultPendingJobsNo, "number of jobs for workers")
	flag.Parse()
}

func main() {
	fmt.Printf("Starting concurrent workers: %d (limit by CPUs no) on %d jobs\n",
		cap(limit),
		*jobsNo)

	// t = make(map[int]int)
	workers := func(l chan<- interface{}, wg *sync.WaitGroup) {
		for i := 0; i <= *jobsNo; i++ {
			i := i

			limit <- struct{}{}
			wg.Add(1)

			go func(x int, w *sync.WaitGroup) {
				defer w.Done()

				// t[i] = rand.Intn(100)
				time.Sleep(1 * time.Second)
				fmt.Printf("Work done: %d\n", i)

				<-limit
			}(i, wg)
		}
	}

	workers(limit, &wg)
	wg.Wait()

	fmt.Println("Finished")
}
