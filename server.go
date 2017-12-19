package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var v string = r.FormValue("hoge")
	fmt.Fprintf(w, "<b>Hello, World :" + v + "</b>")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
