package storage

import (
	"encoding/json"
	"errors"
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
	Status   string `json:"status"`
	Priority string `json:"priority"`
}

var (
	id = -1
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
		Status:   "backlog",
		Priority: "low",
	}
	return a, nil
}

func SaveToFile(newData ToDoList) error {
	allData := LoadDataFromFile()
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

func changeSave(allData []ToDoList) error {
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
	var allData []ToDoList
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Не удалось прочитать файл %s: %v\n", fileName, err)
	}

	isEmpty, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if isEmpty.Size() != 0 {
		err = json.Unmarshal(file, &allData)
		if err != nil {
			log.Fatalf("Ошибка при декодировании JSON из файла %s: %v\n", fileName, err)
		}
	}
	log.Printf("Загружено %d задач из файла %s\n", len(allData), fileName)
	return allData
}

func clearFile() error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func AddToDataBase(numbTasks int) error {
	isEmpty, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if isEmpty.Size() != 0 {
		allDataNew := LoadDataFromFile()
		if len(allDataNew) == 0 {
			err = clearFile()
			if err != nil {
				return err
			}
		}
		id = allDataNew[len(allDataNew)-1].Id
	}
	for i := 0; i < numbTasks; i++ {
		data, err := fill()
		if err != nil {
			log.Println("Ошибка заполнения данных:", err)
			return err
		}

		err = SaveToFile(data)
		if err != nil {
			log.Println("Ошибка сохранения в файл:", err)
			return err
		} else {
			log.Println("Данные успешно сохранены в", fileName, "; Запись:", i+1)
		}
	}
	return nil
}

func DeleteByID(deleteID int) error {
	isEmpty, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if isEmpty.Size() != 0 {
		allDataNew := LoadDataFromFile()
		for i := 0; i < len(allDataNew); i++ {
			if allDataNew[i].Id == deleteID {
				allDataNew = append(allDataNew[:i], allDataNew[i+1:]...)
				err := changeSave(allDataNew)
				if err == nil {
					return nil
				}
			}
		}
	}
	return errors.New("No task to delete by this ID")
}

func EditByID(idX int, editName string, editDeadline string, editInfo string, priority string) error {
	if idX < 0 {
		return errors.New("Imposible")
	}
	isEmpty, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if isEmpty.Size() == 0 {
		NewTask := ToDoList{
			Id:       idX,
			NameNumb: editName,
			Deadline: editDeadline,
			Info:     editInfo,
		}
		fmt.Println(NewTask)
		err := SaveToFile(NewTask)
		if err != nil {
			return err
		}
		return nil
	}
	allDataNew := LoadDataFromFile()
	if len(allDataNew) == 0 {
		err := clearFile()
		if err != nil {
			return err
		}
	}
	cheat := len(allDataNew)
	for i := 0; i < len(allDataNew); i++ {
		if allDataNew[i].Id == idX {
			allDataNew[i].Info = editInfo
			allDataNew[i].Deadline = editDeadline
			allDataNew[i].NameNumb = editName
			allDataNew[i].Id = idX
			allDataNew[i].Priority = priority
			err := clearFile()
			if err != nil {
				return err
			}
			fmt.Println(allDataNew)

			for j := 0; j < cheat; j++ {
				fmt.Println(j, allDataNew[j])
				err := SaveToFile(allDataNew[j])
				if err != nil {
					return err
				}
			}
			return nil
		}
		if i == len(allDataNew)-1 {
			NewTask := ToDoList{
				Id:       idX,
				NameNumb: editName,
				Deadline: editDeadline,
				Info:     editInfo,
				Priority: priority,
			}
			fmt.Println(NewTask)
			err := SaveToFile(NewTask)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func EditStatusByID(idX int, edit_Status string) error {
	if idX < 0 {
		return errors.New("Imposible")
	}
	isEmpty, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if isEmpty.Size() == 0 {
		return errors.New("Empty")
	}
	allDataNew := LoadDataFromFile()
	if len(allDataNew) == 0 {
		return errors.New("Empty")
	}
	cheat := len(allDataNew)
	for i := 0; i < len(allDataNew); i++ {
		if allDataNew[i].Id == idX {
			allDataNew[i].Status = edit_Status
			err := clearFile()
			if err != nil {
				return err
			}

			for j := 0; j < cheat; j++ {
				fmt.Println(j, allDataNew[j])
				err := SaveToFile(allDataNew[j])
				if err != nil {
					return err
				}
			}
			return nil
		}
		if i == len(allDataNew)-1 {
			return errors.New("Nothing to edit")
		}
	}
	return nil
}
