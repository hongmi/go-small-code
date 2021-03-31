package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	mutex := &sync.Mutex{}

	var n uint64 = 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for k := 0; k < 100; k++ {
				mutex.Lock()
				n++
				mutex.Unlock()
			}
		}()

	}

	wg.Wait()
	fmt.Println(n)

}
