package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// PageData ...
// Holds some of the metadata for each page to be served
type PageData struct {
	Title string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := PageData{Title: "Index"}
	t, err := template.ParseFiles("./public/index.html")
	if err != nil {
		fmt.Println("Error parsing template")
	}
	t.Execute(w, p)
}

func main() {
	port := ":4040"
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(port, nil)
}
