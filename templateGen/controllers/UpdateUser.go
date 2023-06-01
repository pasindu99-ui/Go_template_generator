package controllers

import (
	"fmt"
	"os"
	"text/template"
)

type DetailsUpdate struct {
	MongoURL       string
	DatabaseName   string
	CollectionName string
}

func UpdateUser(url, dbName, collectionName string) {

	details := Details{url, dbName, collectionName}
	templatePath := "F:\\Go_template_generator\\templateGen\\templates\\updateTemplate.txt"

	t, err := template.New("updateTemplate.txt").ParseFiles(templatePath)

	if err != nil {
		fmt.Println("Error executing template:", err)
	}

	outputFile := "crud_app/controllers/updateUser.go"
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
