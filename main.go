package main

import (
	"fmt"
	"net/http"
)

func handleFunc(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "<p>Hello people</p>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", nil)
}
