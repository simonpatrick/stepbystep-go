package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Request")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9001", nil)
}
