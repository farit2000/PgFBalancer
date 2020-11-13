package web

import (
	"html/template"
	"log"
	"net/http"
)


type Todo struct {
	Title string
	Done  bool
}


type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}


func StartWeb() {
	tmpl := template.Must(template.ParseFiles("web/views/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "PgFBalancer",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			log.Print(err.Error())
		}
	})
	err := http.ListenAndServe(":5678", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
