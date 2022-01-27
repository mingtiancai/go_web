package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "handle-hello")
}

func world(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "handle-world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	server.ListenAndServe()
}
