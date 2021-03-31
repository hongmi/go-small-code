package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"abc", "bca", "acb"}

	sort.Strings(a)

	ii := []int{2, 3, 4, 5}

	b := sort.IntsAreSorted(ii)

	fmt.Println(a)
	fmt.Println(b)

}
