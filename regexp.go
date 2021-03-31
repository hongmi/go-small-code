package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("[a-z]+", "abc")
	fmt.Println(match)
}
