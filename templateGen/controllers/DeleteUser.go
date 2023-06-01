package controllers

import (
	"fmt"
	"os"
	"text/template"
)

type DetailsDelete struct {
	MongoURL       string
	DatabaseName   string
	CollectionName string
}

func DeleteUser(url, dbName, collectionName string) {

	details := Details{url, dbName, collectionName}
	templatePath := "F:\\Go_template_generator\\templateGen\\templates\\deleteTemplate.txt"

	t, err := template.New("deleteTemplate.txt").ParseFiles(templatePath)

	if err != nil {
		fmt.Println(err)
	}

	outputFile := "crud_app/controllers/deleteUser.go"
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
