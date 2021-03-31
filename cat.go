package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Errorf("Usage: cat filepath")
		return
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("fail to open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("", err)
	}
}
