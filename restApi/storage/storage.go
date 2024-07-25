package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type ToDoList struct {
	Id       int    `json:"id"`
	NameNumb string `json:"name_numb"`
	Deadline string `json:"deadline"`
	Info     string `json:"info"`
}

var (
	allData []ToDoList
	id      = -1
)

const fileName = "berendeev/restApi/storage/storage.json"

func timing() string {
	deadLine := time.Now().Add(24 * time.Hour)
	y := strconv.Itoa(deadLine.Year())
	m := deadLine.Month().String()
	d := strconv.Itoa(deadLine.Day())
	date := d + " " + m + " " + y
	return date
}

func fill() (ToDoList, error) {
	id += 1
	a := ToDoList{
		Id:       id,
		NameNumb: "Task " + strconv.Itoa(id),
		Deadline: timing(),
		Info:     "Nothing",
	}
	return a, nil
}

func SaveToFile(newData ToDoList) error {
	allData = append(allData, newData)

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(allData)
}

func LoadDataFromFile() []ToDoList {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Не удалось прочитать файл %s: %v\n", fileName, err)
	}

	err = json.Unmarshal(file, &allData)
	if err != nil {
		log.Fatalf("Ошибка при декодировании JSON из файла %s: %v\n", fileName, err)
	}

	log.Printf("Загружено %d задач из файла %s\n", len(allData), fileName)
	return allData
}

func StartDataBase() {
	isEmpty, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if isEmpty.Size() != 0 {
		allData = LoadDataFromFile()
		id = allData[len(allData)-1].Id
	}
	for i := 0; i < 10; i++ {
		data, err := fill()
		if err != nil {
			fmt.Println("Ошибка заполнения данных:", err)
			return
		}

		err = SaveToFile(data)
		if err != nil {
			fmt.Println("Ошибка сохранения в файл:", err)
		} else {
			fmt.Println("Данные успешно сохранены в", fileName, "; Запись:", i+1)
		}
	}
}
