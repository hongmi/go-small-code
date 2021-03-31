package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("1.loop i=%d\n", i)
	}

	i := 0
	for i < 10 {
		fmt.Printf("2.loop i=%d\n", i)
		i++
	}

	go func() {
		k := 0
		for {
			fmt.Printf("infinite loop k=%d\n", k)
			k++
			if k > 5 {
				return
			}
		}
	}()

	select {
	case <-time.After(time.Second):
	}
}
