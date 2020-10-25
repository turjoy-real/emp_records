package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

type Emp struct {
	Id int
	Name string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/records", empIndex)
	http.HandleFunc("/records/show", empShow)
	http.HandleFunc("/records/create", empCreate)
	http.HandleFunc("/records/create/process", empCreateProcess)
	http.HandleFunc("/records/create", empUpdateForm)
	http.HandleFunc("/records/update/process", empUpdateProcess)
}

func index() {
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

func empIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	
}
