package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: alert [seconds]")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Errorf("error number of seconds", err)
		return
	}

	timeout := make(chan int, 100)

	go func() {
		for i := n; i >= 0; i-- {
			timeout <- i
			time.Sleep(time.Second)
		}

	}()

	for {
		select {
		case seconds := <-timeout:
			// http://ascii-table.com/ansi-escape-sequences.php
			fmt.Printf("\033[100D\033[K%d", seconds)
			if seconds == 0 {
				goto a
			}
		}
	}

a:
	fmt.Println("\a")
}
