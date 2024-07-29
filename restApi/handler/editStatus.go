package handler

import (
	"advanced-tasks/berendeev/restApi/storage"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func pageEditStatus(w http.ResponseWriter, mb string) {
	htmlContent, _ := os.ReadFile("berendeev/restApi/front-end/editStatus" + mb + ".html")
	fmt.Fprintf(w, string(htmlContent))
	return
}

func EditStatusHandler() {
	http.HandleFunc("/editStatus", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		editStatus := r.FormValue("edit_Status")
		IdChangeStatusMB := r.FormValue("IdChangeStatus")

		IdChangeStatus, err := strconv.Atoi(IdChangeStatusMB)
		if err != nil {
			pageEditStatus(w, "Error")
			log.Print(err)
			return
		}

		err = storage.EditStatusByID(IdChangeStatus, editStatus)
		if err != nil {
			pageEditStatus(w, "Error")
			return
		}
		pageEditStatus(w, "")
	})
}
