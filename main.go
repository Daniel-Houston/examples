package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

const schema = `
{
	"type":"object",
	"properties":{
		"id":{
			"type":"string"
		},
		"description":{
			"type":"string"
		}
	}
}
`

func main() {
	file := "file.json"
	canonical, err := filepath.Abs(file)
	if err != nil {
		log.Fatal(err)
	}

	schemaLoader := gojsonschema.NewStringLoader(schema)
	documentLoader := gojsonschema.NewReferenceLoader("file://" + canonical)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		log.Fatal(err)
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
