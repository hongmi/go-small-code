package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 10)

	for i := 0; i < 10; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(time.Second)

	for req := range requests {
		<-limiter.C
		fmt.Println(req, time.Now())
	}

}
