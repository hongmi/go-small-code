package main

import (
	"fmt"

	"time"
)

func main() {
	t := time.NewTicker(2 * time.Second)

	done := make(chan struct{})

	go func() {

		for {

			select {
			case <-done:
				return
			case tk := <-t.C:
				fmt.Println("tick at", tk)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	done <- struct{}{}

}
