package main

import (
	"fmt"
	"sync"
	//	"time"
)

func main() {
	/*
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

		f := func() {
			in := gen(numbers...)
			c1 := sq(in)
			c2 := sq(in)
			c3 := sq(in)

			for n := range merge(c1, c2, c3) {
				fmt.Printf("%d ", n)
			}
			fmt.Println()
		}

		go f()
		time.Sleep(time.Second)
		go f()
		time.Sleep(time.Second)
		go f()
		time.Sleep(time.Second)
		go f()
		time.Sleep(time.Second)

		ci := gen(numbers...)
		fmt.Println(<-ci)
	*/
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from output.
	done := make(chan int, 2)
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// Tell the remaining senders we're leaving.
	//done <- 1
	//	done <- 2

}

func gen(numbers ...int) <-chan int {

	c := make(chan int)

	go func() {
		for _, i := range numbers {
			c <- i
		}

		close(c)
	}()
	return c
}

func sq(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		for i := range in {
			out <- i
		}

		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	wg.Add(len(cs))

	out := make(chan int)

	for _, c := range cs {
		go func(ch <-chan int) {
			for i := range ch {
				out <- i
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
