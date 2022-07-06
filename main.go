package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	res := hello(r.URL.RawQuery)
	fmt.Fprintf(w, "%s\n", res)
}

func hello(s string) string {
	return "Hello, " + s
}
