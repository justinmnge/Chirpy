package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("."))
	mux.Handle("/", fs)
	srv := http.Server{Addr: ":8080", Handler: mux}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
