package main

import (
	"fmt"
	"time"
)

type tt struct {
}

func (ttt tt) F() {
}

func main() {

	v := "ss"

	go func(v interface{}) {
		switch t := v.(type) {
		case bool:
			fmt.Println("bool type")
		default:
			fmt.Printf("unkown type:%T\n", t)
		}
	}(v)

	time.Sleep(time.Second)
}
