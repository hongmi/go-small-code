package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", hello)

	fmt.Println("http://locahost:8090/hello")
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
