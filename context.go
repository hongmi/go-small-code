package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("http://localhost:8090/hello")
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()

	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Hello\n")
	}
}
