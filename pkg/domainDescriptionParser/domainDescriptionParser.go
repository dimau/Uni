package domainDescriptionParser

import (
	"encoding/json"
	"log"
)

type ValueType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Attribute struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Database  string    `json:"database"`
	ValueType ValueType `json:"value-type"`
}

type Class struct {
	Id         string      `json:"id"`
	Name       string      `json:"name"`
	Database   string      `json:"database"`
	Attributes []Attribute `json:"attributes"`
}

type Domain struct {
	Classes []Class `json:"classes"`
}

func ParseJsonDomainDescription(domain []byte) *Domain {
	var data = &Domain{}
	err := json.Unmarshal(domain, data)
	if err != nil {
		log.Fatalln("Error unmarshalling", err.Error())
	}
	return data
}
