package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))

	mux.Handle("/", fs)
	http.ListenAndServe(":3000", mux)
}
