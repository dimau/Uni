package appSourceCodeCreator

import (
	"fmt"
	"github.com/dimau/Uni/pkg/domainDescriptionParser"
)

func CreateAppSourceCode(jsonDomainDescription *domainDescriptionParser.Domain) *[]byte {
	sourceFileContent := []byte("package main\n\n")

	for _, class := range jsonDomainDescription.Classes {
		sourceFileContent = append(sourceFileContent, []byte(fmt.Sprintf("type %v struct {\n", class.Name))...)
		for _, attribute := range class.Attributes {
			sourceFileContent = append(sourceFileContent, []byte(fmt.Sprintf("    %v %v\n", attribute.Name, attribute.ValueType.Name))...)
		}
		sourceFileContent = append(sourceFileContent, []byte(fmt.Sprintf("}\n\n"))...)
	}

	return &sourceFileContent
}
