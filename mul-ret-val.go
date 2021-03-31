package main

import (
	"fmt"
)

func foo() (int, bool) {
	return 1, false
}

func main() {

	fmt.Println(foo())

}
