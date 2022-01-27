package main

import (
	"fmt"
	"net/http"
)

type HelloHandle struct{}

func (h HelloHandle) ServeHTTP(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "hello!")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, q *http.Request) {
		fmt.Printf("Handle called - %T\n", h)
		h.ServeHTTP(w, q)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, q *http.Request) {
		h.ServeHTTP(w, q)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := HelloHandle{}
	http.Handle("/hello/", protect(log(hello)))
	server.ListenAndServe()
}
