package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	message, ok := os.LookupEnv("APP_NAME")
	if !ok {
		message = "Sample Web"
	}

	h := func(w http.ResponseWriter, _ *http.Request) {
		for i := 0; i < 10000000000; i++ {
			// do nothing
		}
		io.WriteString(w, message+"\n")
	}
	http.HandleFunc("/", h)
	http.ListenAndServe(":8080", nil)
}
