package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello")
	case "/varun":
		fmt.Fprint(w, "Varun")
	default:
		fmt.Fprint(w, "Error")
	}
	fmt.Printf("Handling request of type %s\n", r.Method)
}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain") // text/html is default
	fmt.Fprint(w, "<h1>Hello</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Timeout")
	time.Sleep(2 * time.Second)
	fmt.Printf("Timeout")
}

func main() {
	// http.HandleFunc("/", helloWorldPage)
	http.HandleFunc("/", htmlVsPlain)
	http.HandleFunc("/timeout", timeout)
	// http.ListenAndServe("", nil)
	server := http.Server{
		Addr:         "",
		Handler:      nil,
		WriteTimeout: 1000,
		ReadTimeout:  1000,
	}
	server.ListenAndServe()
}
