package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("html/")))

	log.Println("Start server from localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
