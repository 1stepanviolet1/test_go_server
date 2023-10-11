package main

import (
	"log"
	"net/http"
	"os"
)


func send_index(w http.ResponseWriter, r *http.Request) {
	html_index := get_index()
	w.Write(html_index)

}


func send_styles(w http.ResponseWriter, r *http.Request) {
	styles := get_styles()
	w.Write(styles)

}


func main() {
	http.HandleFunc("/", send_index)
	http.HandleFunc("/style.css", send_styles)

	log.Println("Start server from localhost")
	log.Fatal(http.ListenAndServe("localhost:80", nil))

}


func get_index() []byte {
	html, err := os.ReadFile("./html/index.html")
	
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return html

}


func get_styles() []byte {
	styles, err := os.ReadFile("./html/style.css")
	
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return styles

}
