package main

import (
	"fmt"
	"net/http"
)

func handler1(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path)
}

func handle2(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(write, "ppppp")
}

func main() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/p.", handle2)
	http.ListenAndServe(":8080", nil)
}
