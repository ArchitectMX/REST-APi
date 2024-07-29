package handler

import (
	"advanced-tasks/berendeev/restApi/storage"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var taskNumbPost = 0

func errPagePost(w http.ResponseWriter) {
	htmlContent, _ := os.ReadFile("berendeev/restApi/front-end/postError.html")
	fmt.Fprintf(w, string(htmlContent))
	return
}

func PostHandler() {
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		tasksNumber := r.FormValue("task_numb")

		var err error
		taskNumbPost, err = strconv.Atoi(tasksNumber)
		if err != nil || taskNumbPost < 0 {
			errPagePost(w)
			return
		}

		err = storage.AddToDataBase(taskNumbPost)
		if err != nil {
			errPagePost(w)
			return
		}
		htmlContent, err := os.ReadFile("berendeev/restApi/front-end/post.html")
		if err != nil {
			errPagePost(w)
			return
		}
		fmt.Fprintf(w, string(htmlContent))
	})
}
