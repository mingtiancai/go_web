package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// fmt.Fprintf(w, r.PostForm["hello"][0])
	// fmt.Fprintf(w, "\n")
	// fmt.Fprintln(w, r.Form)
	// r.ParseMultipartForm(1024)
	// fmt.Fprintln(w, r.MultipartForm)
	// fmt.Fprintf(w, r.MultipartForm.Value["hello"][0])
	fmt.Fprintf(w, r.FormValue("hello"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
