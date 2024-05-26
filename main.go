package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", helloWorldPage)
	http.ListenAndServe("", nil) // empty "" means it will use default 80 as port: open localhost:80
}
