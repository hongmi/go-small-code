package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup

	var n uint64 = 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for k := 0; k < 100; k++ {
				atomic.AddUint64(&n, 1)
			}
		}()

	}

	wg.Wait()
	fmt.Println(n)
}
