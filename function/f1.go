package main

import (
	"fmt"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(greeting))
}
