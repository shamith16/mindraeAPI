package utils

import (
	"log"
	"os"
)

func WriteToFile(filename string, data []byte, path string, operationType string) (err error) {
	var permission int

	switch {
	case operationType == "append":
		permission = os.O_WRONLY | os.O_TRUNC | os.O_CREATE | os.O_APPEND
	case operationType == "new":
		permission = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
	}

	file, err := os.OpenFile(
		path+filename,
		permission,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	// Write bytes to file
	byteSlice := data
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Wrote %d bytes.\n", bytesWritten)

	return
}
