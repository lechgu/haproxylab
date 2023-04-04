package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", authenticate)
	http.ListenAndServe(":6060", nil)
}

func authenticate(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
