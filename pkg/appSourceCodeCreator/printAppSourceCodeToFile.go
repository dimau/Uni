package appSourceCodeCreator

import (
	"log"
	"os"
)

func PrintAppSourceCodeToFile(sourceFileContent *[]byte) {
	// Create or (if already created) clean current file and open for writing
	file, err := os.Create("./resultApplication/application.go")
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Closing file after all writes operations done
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	// Writes all bytes to text file with source code for the App
	_, err = file.Write(*sourceFileContent)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
