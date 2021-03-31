package main

import "fmt"

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	i := c()
	fmt.Println(i)
}
