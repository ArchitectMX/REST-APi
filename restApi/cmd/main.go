package main

import (
	"advanced-tasks/berendeev/restApi/front-end"
	"fmt"
	"net/http"
)

func main() {
	front_end.PageGenerator()

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}
