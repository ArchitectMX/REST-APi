package fileWork

import "os"

func parsing() {
	file, err := os.Open("C:\\Users\\tberendeev\\Desktop\\forJson\\TEST.json")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

}
