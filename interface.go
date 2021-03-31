package main

import (
	"fmt"
	"sort"
)

type S1 []string

func (s S1) First() string {
	return s[0]
}

func (s S1) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s S1) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s S1) Len() int {
	return len(s)
}

func main() {
	a := []string{"a", "b", "c"}

	s := S1(a)

	sort.Sort(s)

	fmt.Println(s)
}
