package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Title     string
	FirstName string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

func main() {

	http.HandleFunc("/", idx)
	http.HandleFunc("/about", abt)
	http.HandleFunc("/contact", cntct)
	http.HandleFunc("/signin", sgnin)
	http.Handle("/favicon.ico", http.NotFoundHandler()) //to handle errors
	http.ListenAndServe(":3000", nil)
}

func idx(w http.ResponseWriter, r *http.Request) {

	pd := pageData{
		Title: "Index Page",
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)
	if err != nil {
		log.Println("Logged", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// fmt.Println("We got into index")
	// fmt.Println(r.URL.Path)
}

func abt(w http.ResponseWriter, r *http.Request) {

	pd := pageData{
		Title: "About Page",
	}

	err := tpl.ExecuteTemplate(w, "about.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func cntct(w http.ResponseWriter, r *http.Request) {

	pd := pageData{
		Title: "Contact Page",
	}

	err := tpl.ExecuteTemplate(w, "contact.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func sgnin(w http.ResponseWriter, r *http.Request) {

	pd := pageData{
		Title: "Sign In Page",
	}

	if r.Method == http.MethodPost {
		first := r.FormValue("fname")
		pd.FirstName = first
	}

	err := tpl.ExecuteTemplate(w, "signin.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
