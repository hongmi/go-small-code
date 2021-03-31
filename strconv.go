package main

import (
	"fmt"
	"strconv"
)

func main() {

	p := fmt.Println

	b, _ := strconv.ParseBool("false")

	p(b)

	i, _ := strconv.ParseInt("889", 16, 64)
	p(i)
}
