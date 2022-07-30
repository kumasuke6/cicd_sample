package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", echo)
	http.HandleFunc("/html", html)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	res := hello(r.URL.RawQuery)
	fmt.Fprintf(w, "%s\n", res)
}

func html(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func hello(s string) string {
	return "Hello, " + s
}
