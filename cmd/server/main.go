package main

import (
	"html/template"
	"net/http"

	page "github.com/hnefatafl01/indy-budget/domain/page"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../../tmpl/*.html"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../../assets"))))
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)

}

func renderTemplate(w http.ResponseWriter, n string, d interface{}) {
	err := tpl.ExecuteTemplate(w, n, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	p1 := page.Page{
		Title: "Home",
	}
	// w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "index.html", p1)
}
