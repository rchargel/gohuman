package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const indexContent template.HTML = template.HTML("<h1>Testing</h1>")
const cssContent string = `body {
  font-family: Arial, Helvetica, sans-serif;
  background-color: #fff;
  color: #000;
}`

// IndexController the index controller.
type IndexController struct{}

// CSSHandler renders the css file for this application.
func (i IndexController) CSSHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/css")
	fmt.Fprintf(w, cssContent)
}

// Handler the handler for the index controller.
func (i IndexController) Handler(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Title: "Go-Human: We Love Humans!",
		Body:  indexContent,
	}
	p.Render(w)
}
