package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	val   uint64
	respC chan uint64
}

func main() {

	var wg sync.WaitGroup

	var n uint64 = 0

	msgs := make(chan Msg)

	go func() {
		for {
			select {
			case msg := <-msgs:
				n++
				msg.respC <- n
			}
		}
	}()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for k := 0; k < 100; k++ {
				msg := Msg{
					val:   n,
					respC: make(chan uint64),
				}
				msgs <- msg
				<-msg.respC
			}
		}()

	}

	wg.Wait()
	fmt.Println(n)

}
