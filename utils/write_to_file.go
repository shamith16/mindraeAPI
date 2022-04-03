package utils

import (
	"log"
	"os"
	"path/filepath"
)

func WriteToFile(filename string, data interface{}, path string, operationType string) (err error) {
	var permission int
	var permMode int

	switch {
	case operationType == "append":
		permission = os.O_APPEND | os.O_CREATE | os.O_WRONLY
		permMode = 0644
		break
	case operationType == "new":
		permission = os.O_WRONLY | os.O_TRUNC | os.O_CREATE
		permMode = 0644
		break
	}

	//Get the base file dir
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("error msg", err)
	}

	//Create output path
	outPath := filepath.Join(cwd, path)

	//Create dir output using above code
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		os.Mkdir(outPath, 0755)
	}

	file, err := os.OpenFile(
		path+filename,
		permission,
		os.FileMode(permMode),
	)
	//s := fmt.Sprintf("Currently writing/appending to %s with permission %d and permMode is %d", path+filename, permission, permMode)
	//fmt.Println(s)
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
	var bytesWritten int

	switch data.(type) {
	case []string:
		{
			stringList := data.([]string)
			for _, s := range stringList {
				bytesWritten, err = file.WriteString(s)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	case []byte:
		{
			dataByte := data.([]byte)
			bytesWritten, err = file.Write(dataByte)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Printf("Wrote %d bytes.\n", bytesWritten)

	return
}
