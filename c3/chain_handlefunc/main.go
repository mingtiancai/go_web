package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

func hello(w http.ResponseWriter, q *http.Request) {
	respInfo := "hello" + time.Now().String()
	fmt.Fprintf(w, respInfo)
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println(time.Now())
		fmt.Println("Handle function called - ", name)
		time.Sleep(2 * time.Second)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}
