package handler

import (
	"advanced-tasks/berendeev/restApi/storage"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func errPageDelete(w http.ResponseWriter) {
	htmlContent, _ := os.ReadFile("berendeev/restApi/front-end/deleteError.html")
	if _, err := fmt.Fprintf(w, string(htmlContent)); err != nil {
		return
	}
	return
}

var tasksNumbDel = 0

func DeleteHandler() {
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		tasksNumber := r.FormValue("task_ID")
		if tasksNumber == "-0" {
			errPageDelete(w)
			return
		}

		var err error
		tasksNumbDel, err = strconv.Atoi(tasksNumber)
		if err != nil || tasksNumbDel < 0 {
			errPageDelete(w)
			return
		}

		err = storage.DeleteByID(tasksNumbDel)
		if err != nil {
			errPageDelete(w)
			return
		}
		htmlContent, err := os.ReadFile("berendeev/restApi/front-end/delete.html")
		if err != nil {
			errPageDelete(w)
			return
		}
		fmt.Fprintf(w, string(htmlContent))
	})
}
