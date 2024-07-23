package main

import (
	"log"
	"net/http"

	"github.com/bek-shoyatbek/beksblog/handlers"
	"github.com/bek-shoyatbek/beksblog/lib"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /pages/{slug}", handlers.PostHandler(lib.FileReader{}))

	err := http.ListenAndServe(":3030", mux)
	if err != nil {
		log.Fatal(err)
	}

}
