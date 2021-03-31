package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	p := fmt.Println

	rand.Seed(time.Now().UnixNano())

	p(rand.Intn(1000))

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	p(r.Intn(100))

}
