package handler

import (
	"advanced-tasks/berendeev/restApi/storage"
	"fmt"
	"log"
	"net/http"
	"os"
)

const fileName = "berendeev/restApi/storage/storage.json"

func HomePage() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := os.ReadFile("berendeev/restApi/front-end/view.html")
		if err != nil {
			http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(htmlContent))

		isEmpty, err := os.Stat(fileName)
		if err != nil {
			log.Fatal(err)
		}
		if isEmpty.Size() != 0 {
			tasks := storage.LoadDataFromFile()
			for i := 0; i < len(tasks); i++ {
				fmt.Fprintf(w, `
<head>
    <style>
        .containerX {
            margin-left: 20px;
            text-align: left; / *  Центрирование текста  * /
        }
    </style>
</head>
<body>
<div class="containerX">
    <div style="margin-left: 45px">
		<pre>
        <div class="message">%v: </div>
        <div class="message">       ID: %v</div>
        <div class="message">       Deadline: %v</div>
        <div class="message">       Info: %v</div>
        <div class="message">       Status: %v</div>
        <div class="message">       Priority: %v</div>
		</pre>
    </div>
</div>
</body>
				        		`, tasks[i].NameNumb, tasks[i].Id, tasks[i].Deadline, tasks[i].Info, tasks[i].Status, tasks[i].Priority)
				if err != nil {
					return
				}
			}
		}
	})
}
