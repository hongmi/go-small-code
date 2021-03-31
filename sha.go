package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	h := sha1.New()

	h.Write([]byte("abc"))
	fmt.Printf("%x\n", h.Sum(nil))
	fmt.Printf("%x\n", h.Sum([]byte("ab")))
	fmt.Printf("%x\n", sha1.Sum([]byte("abc")))

}
