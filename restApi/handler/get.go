package handler

import (
	"advanced-tasks/berendeev/restApi/storage"
	"html/template"
	"net/http"
)

func GetHandler() {
	http.HandleFunc("/get", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content Type", "text/html")

		storage.StartDataBase()

		tasks := storage.LoadDataFromFile()
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, tasks)

	})

	http.ListenAndServe(":8080/get", nil)
}

const docList = `
<ul >
   {{range .}}
   <li>{{.Title}}</li>
   {{end}}
</ul>
`

const doc = `
<!DOCTYPE html>
<html>
   <head><title>{{.Title}}</title></head>
   <body>
       <h1>Hello Templates</h1>
       {{template "List" .TopMovies}}
   </body>
</html>
`
