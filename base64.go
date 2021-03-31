package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "abcdef"
	ss := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(ss)

	b, _ := base64.StdEncoding.DecodeString(ss)
	fmt.Printf(string(b))
}
