package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("defer")

	os.Exit(3)

}
