package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("/templates/*.html"))
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

}

func about(w http.ResponseWriter, req *http.Request) {

}

func apply(w http.ResponseWriter, req *http.Request) {

}

func contact(w http.ResponseWriter, req *http.Request) {

}
