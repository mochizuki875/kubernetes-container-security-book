package main

import (
	"io"
	"net/http"
)

func main() {
	h := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello World!!\n")
	}
	http.HandleFunc("/", h)
	http.ListenAndServe(":8080", nil)
}
