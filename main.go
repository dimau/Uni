package main

import (
	"github.com/dimau/Uni/pkg/appSourceCodeCreator"
	"github.com/dimau/Uni/pkg/domainDescriptionParser"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Read marshalled text of JSON Domain Description from the file
	file, err := os.Open("./resultApplication/textDomainDescription.json")
	if err != nil {
		log.Fatalln("Can't open text file with JSON domain description " + err.Error())
	}

	// Closing file after all read operations done
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	// Reading file with domain description
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("Can't read from text file with JSON domain description " + err.Error())
	}

	// Handle JSON domain description and make a source code for an Application
	jsonDomainDescription := domainDescriptionParser.ParseJsonDomainDescription(data)
	sourceFileContent := appSourceCodeCreator.CreateAppSourceCode(jsonDomainDescription)
	appSourceCodeCreator.PrintAppSourceCodeToFile(sourceFileContent)
}
