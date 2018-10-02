package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// PageData ...
// Holds some of the metadata for each page to be served
type PageData struct {
	Title string
}

func (t ToDo) String() string {
	return fmt.Sprintf("{Content:%s, Done:%v, Created:%s, Due:%s}", t.content, t.completed, t.created, t.due)
}

// ToDo ...
// Holds the content of a submitted todo
type ToDo struct {
	content   string
	completed bool
	created   time.Time
	due       time.Time
}

func newTodoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse the form that was submitted
	due := time.Now().Add(time.Hour * 24)

	// create an instance of the todo using the form data
	todo := ToDo{
		content:   r.Form["content"][0],
		completed: false,
		created:   time.Now(),
		due:       due,
	}
	// just return the "somewhat stringified" json back to the user
	fmt.Fprintf(w, todo.String())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := PageData{Title: "To Do's"}
	t, err := template.ParseFiles("./public/index.html")
	if err != nil {
		fmt.Println("Error parsing template")
	}
	t.Execute(w, p)
}

func main() {
	port := ":4040"
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/todo", newTodoHandler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Couldn't Listen:", err)
	}
}
