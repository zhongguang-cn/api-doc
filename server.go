package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./docs"))
	http.Handle("/", fs)

	port := ":8081"
	log.Printf("Start to listen on %v", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
