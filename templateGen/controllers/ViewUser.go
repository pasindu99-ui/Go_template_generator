package controllers

import (
	"fmt"
	"os"
	"text/template"
)

type DetailsView struct {
	MongoURL       string
	DatabaseName   string
	CollectionName string
}

func ViewUser(url, dbName, collectionName string) {

	details := Details{url, dbName, collectionName}
	templatePath := "F:\\Go_template_generator\\templateGen\\templates\\viewTemplate.txt"

	t, err := template.New("viewTemplate.txt").ParseFiles(templatePath)

	if err != nil {
		fmt.Println(err)
	}

	outputFile := "crud_app/controllers/viewUser.go"
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
