package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)

	fmt.Printf("Starting web server at port 5005")
	if err := http.ListenAndServe(":5005", nil); err != nil {
		log.Fatal(err)
	}
}
