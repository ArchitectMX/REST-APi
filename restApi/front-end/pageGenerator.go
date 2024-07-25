package front_end

import (
	"advanced-tasks/berendeev/restApi/storage"
	"fmt"
	"net/http"
	"os"
)

func readStage(pathPage string, pathCode string) {
	http.HandleFunc(pathPage, func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := os.ReadFile("berendeev/restApi/front-end/" + pathCode + ".html")
		if err != nil {
			http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(htmlContent))
	})
}

func readStageGet(pathPage string, pathCode string) {
	http.HandleFunc(pathPage, func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := os.ReadFile("berendeev/restApi/front-end/" + pathCode + ".html")
		if err != nil {
			http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(htmlContent))
		storage.StartDataBase()

		tasks := storage.LoadDataFromFile()
		for i := 0; i < len(tasks); i++ {
			fmt.Fprint(w, tasks[i])
		}

	})
}

func PageGenerator() {
	readStage("/", "view")

	readStage("/plug", "plug")

	readStageGet("/get", "get")
}
