package main

import (
	"fmt"
	"net/http"
)

func main() {

	resp, _ := http.Get("http://www.baidu.com")
	defer resp.Body.Close()

	fmt.Println(resp.Status)

}
