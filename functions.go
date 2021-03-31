package main

import (
	"fmt"
)

func addnf(n int) func(int) int {
	return func(k int) int {
		return k + n
	}
}

func main() {

	f := addnf(5)

	fmt.Println(f(4), f(6))
}
