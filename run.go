package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Server on port 8080")
	http.ListenAndServe(":8080", nil)
}
