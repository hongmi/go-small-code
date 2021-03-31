package main

import "fmt"

func main() {

	m := map[int]int{1: 1, 11: 2}

	m[11] = 1
	fmt.Print(m[100])
	a, ok := m[1111]

	fmt.Println(a, ok)

}
