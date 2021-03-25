package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json: "Title"`
	Desc    string `json: "desc"`
	Content string `json: "Content`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "test title", Desc: "ioahfdoih", Content: "tyteugfuhfhofi"},
	}

	fmt.Println("Endpoint Hit: All articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Post endpoint worked")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint Hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", allArticles).Methods("POST")
	// myRouter.HandleFunc("/articles", allArticles).Methods("DELETE")
	// myRouter.HandleFunc("/articles", allArticles).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	handleRequests()
}
