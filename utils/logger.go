package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Logger(str string, routeName string) {

	year, month, day := time.Now().Date()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("error msg", err)
	}

	outPath := filepath.Join(cwd, "/log")

	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		os.Mkdir(outPath, 0755)
	}

	filename := fmt.Sprintf("log-[%d-%d-%d]-%s.txt", day, month, year, routeName)

	file, err := os.OpenFile("log/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	file.WriteString(str)

}
