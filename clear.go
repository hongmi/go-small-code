package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	const col = 30
	// Clear the screen by printing \x0c.
	bar := fmt.Sprintf("\033[100D\033[K[%%-%vs]", col)
	for i := 0; i < col; i++ {
		fmt.Printf(bar, strings.Repeat("=", i)+">")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf(bar+" Done!", strings.Repeat("=", col))
}
