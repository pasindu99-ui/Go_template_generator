package controllers

import (
	"fmt"
	"os"
	"text/template"
)

type Details struct {
	MongoURL       string
	DatabaseName   string
	CollectionName string
}

func InsertUser(url, dbName, collectionName string) {
	details := Details{url, dbName, collectionName}
	templatePath := "F:\\Go_template_generator\\templateGen\\templates\\insertTemplate.txt"

	t, err := template.New("insertTemplate.txt").ParseFiles(templatePath)

	if err != nil {
		fmt.Println("Error parsing template:", err)
	}

	outputFile := "crud_app/controllers/insertUser.go"
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating file:", err)

	}
	defer file.Close()

	err = t.Execute(file, details)
	if err != nil {
		fmt.Println("Error executing template:", err)

	}

}
