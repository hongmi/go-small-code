package main

import (
	"fmt"
	"time"
)

func gn(f func(), n int) func() {
	for i := 0; i < n; i++ {
		f = func() { go f() }
	}
	return f
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go gn(func() { c1 <- "hello" }, 0)()
	go gn(func() { c2 <- "world" }, 10000000000)()

	go func() {
		select {
		case msg2 := <-c2:
			fmt.Println(msg2)
		case msg1 := <-c1:
			fmt.Println(msg1)
		}
	}()

	time.Sleep(time.Second)
}
