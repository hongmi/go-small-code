package main

import (
	"fmt"
)

func main() {

	s := []string{"a", "b", "c"}

	s1 := Filter(s, func(e string) bool {
		return e != "b"
	})

	fmt.Println(s1)
}

func Filter(s []string, f func(string) bool) []string {
	res := []string{}
	for _, e := range s {
		if f(e) {
			res = append(res, e)
		}
	}
	return res
}
