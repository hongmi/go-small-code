package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(2 * time.Second)

	go func() {
		<-t.C
	}()

	t.Stop()

	fmt.Println("end")
}
