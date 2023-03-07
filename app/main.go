package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func index(w http.ResponseWriter, req *http.Request) {
	val := os.Getenv("APP_ID")
	if val == "" {
		val = "unknown"
	}
	fmt.Fprintf(w, "hello from app: %s\n", val)
}

func main() {
	ports := os.Getenv("PORT")
	if ports == "" {
		ports = "80"
	}
	var port int64
	var err error
	if port, err = strconv.ParseInt(ports, 10, 32); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalln(err)
	}
}
