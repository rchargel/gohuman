package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Server")
	http.ListenAndServe(":8080", nil)
}
