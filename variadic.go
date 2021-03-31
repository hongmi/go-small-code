package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(numbers ...string) int {
	var res int

	for _, n := range numbers {
		i, _ := strconv.Atoi(n)
		res += i
	}
	return res
}

func main() {

	fmt.Println(sum(os.Args[1:]...))

}
