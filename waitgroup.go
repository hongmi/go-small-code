package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	defer wg.Wait()

	jobs := make(chan int)
	defer close(jobs)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(jobs, &wg)
	}

	jobs <- 1
	jobs <- 2
	jobs <- 3
}

func worker(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	n := <-jobs

	fmt.Println(n)

}
