package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const pendingJobsNo = 50

func main() {
	var wg sync.WaitGroup
	limit := make(chan interface{}, runtime.NumCPU())

	fmt.Printf("Starting concurrent workers: %d (limit by CPUs no)\n", cap(limit))

	workers := func(l chan<- interface{}, wg *sync.WaitGroup) {
		for i := 0; i <= pendingJobsNo; i++ {
			i := i

			limit <- struct{}{}
			wg.Add(1)

			go func(x int, w *sync.WaitGroup) {
				defer w.Done()

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
