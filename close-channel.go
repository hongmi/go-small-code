package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {

		for {
			i, ok := <-ch
			if ok {
				fmt.Printf("channel return ok = %t\n", ok)
				fmt.Printf("get value from channel: %d\n", i)
			} else {
				fmt.Printf("channel return ok = %t\n", ok)
				break
			}
		}

	}()

	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	time.Sleep(time.Second)

}
