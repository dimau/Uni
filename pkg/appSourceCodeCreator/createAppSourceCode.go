package appSourceCodeCreator

import (
	"fmt"
	"github.com/dimau/Uni/pkg/domainDescriptionParser"
	"log"
)

var valueTypes = map[string]string{
	"uni.org/value-types/string": "string",
	"uni.org/value-types/phone":  "string",
	"uni.org/value-types/email":  "string",
}

func getGolangValueTypeById(valueTypeId string) string {
	val, ok := valueTypes[valueTypeId]
	if ok == false {
		log.Fatalln("Incorrect value type ", valueTypeId)
	}
	return val
}

func CreateAppSourceCode(jsonDomainDescription *domainDescriptionParser.Domain) *[]byte {
	sourceFileContent := []byte("package main\n\n")

	for _, class := range jsonDomainDescription.Classes {
		sourceFileContent = append(sourceFileContent, []byte(fmt.Sprintf("type %v struct {\n", class.Name))...)
		for _, attribute := range class.Attributes {
			sourceFileContent = append(sourceFileContent, []byte(fmt.Sprintf("    %v %v\n", attribute.Name, getGolangValueTypeById(attribute.ValueType)))...)
		}
		sourceFileContent = append(sourceFileContent, []byte(fmt.Sprintf("}\n\n"))...)
	}

	return &sourceFileContent
}
