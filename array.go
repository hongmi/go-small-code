package main

import (
	"fmt"
)

func main() {

	var a [4]int = [...]int{1, 2, 3, 4}

	for i, e := range a {

		fmt.Printf("a[%d]=%d\n", i, e)
	}

}
