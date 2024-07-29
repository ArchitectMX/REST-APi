package handler

import (
	"advanced-tasks/berendeev/restApi/storage"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type ToDoList struct {
	Id       int    `json:"id"`
	NameNumb string `json:"name_numb"`
	Deadline string `json:"deadline"`
	Info     string `json:"info"`
	Status   string `json:"status"`
	Priority string `json:"priority"`
}

func pageEdit(w http.ResponseWriter, mb string) {
	htmlContent, _ := os.ReadFile("berendeev/restApi/front-end/edit" + mb + ".html")
	fmt.Fprintf(w, string(htmlContent))
	return
}

func EditHandler() {
	http.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		editIDMay := "-1"
		editName := "Task -1"
		editDeadline := "29 July 2024"
		editInfo := "Nothing"
		if r.FormValue("edit_ID") != "" {
			editIDMay = r.FormValue("edit_ID")
		}
		if r.FormValue("edit_Name") != "" {
			editName = r.FormValue("edit_Name")
		}
		if r.FormValue("edit_Deadline") != "" {
			editDeadline = r.FormValue("edit_Deadline")
		}
		if r.FormValue("edit_Info") != "" {
			editInfo = r.FormValue("edit_Info")
		}
		priority := r.FormValue("priority")

		editId, err := strconv.Atoi(editIDMay)
		if err != nil || taskNumbPost < 0 {
			pageEdit(w, "Error")
			return
		}

		err = storage.EditByID(editId, editName, editDeadline, editInfo, priority)
		if err != nil {
			pageEdit(w, "Error")
			return
		}
		pageEdit(w, "")
	})
}
