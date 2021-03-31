package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	s, _ := json.Marshal(false)
	fmt.Println(string(s))

	var b bool
	json.Unmarshal(s, &b)
	fmt.Println(b)

	s1, _ := json.Marshal(1)
	fmt.Println(string(s1))

}
