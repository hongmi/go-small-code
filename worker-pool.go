package main

import (
	"fmt"
	"time"
)

func main() {

	jobs := make(chan int)
	defer close(jobs)
	for i := 0; i < 3; i++ {
		go worker(jobs)
	}

	jobs <- 1
	jobs <- 2
	jobs <- 3
	time.Sleep(time.Second)
}

func worker(jobs <-chan int) {

	n := <-jobs

	fmt.Println(n)

}
