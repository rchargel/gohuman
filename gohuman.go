package main

import (
	"log"
	"net/http"
)

func init() {
	log.Println("Initializing handlers..")
	http.HandleFunc("/main.css", IndexController{}.CSSHandler)
	http.HandleFunc("/", IndexController{}.Handler)
}
