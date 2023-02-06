package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		panic(err.Error())
	}
}
