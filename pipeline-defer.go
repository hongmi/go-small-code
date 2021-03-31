package main

import (
	"fmt"
	"sync"
	//	"time"
)

func main() {
	in := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	// Distribute the sq work across two goroutines that both read from in.

	// Consume the first value from output.
	done := make(chan bool)
	defer close(done)

	c1 := sq(done, in)
	c2 := sq(done, in)

	out := merge(done, c1, c2)
	//fmt.Println(<-out) // 4 or 9

	for n := range out {
		fmt.Printf("%d ", n)
	}

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

func sq(done <-chan bool, in <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			select {
			case out <- i * i:
			case <-done:
				return
			}
		}
	}()

	return out
}

func merge(done <-chan bool, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	wg.Add(len(cs))

	out := make(chan int)

	for _, c := range cs {
		go func(ch <-chan int) {
			defer wg.Done()
			for i := range ch {
				select {
				case out <- i:
				case <-done:
					return
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
