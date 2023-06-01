package controllers

import (
	"fmt"
	"os"
	"text/template"
)

func ServerFile() {
	templatePath := "F:\\Go_template_generator\\templateGen\\templates\\serverTemplate.txt"

	t, err := template.New("serverTemplate.txt").ParseFiles(templatePath)

	if err != nil {
		fmt.Println("Error parsing template:", err)
	}

	outputFile := "crud_app/server.go"
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating file:", err)

	}
	defer file.Close()

	err = t.Execute(file, nil)
	if err != nil {
		fmt.Println("Error executing template:", err)

	}

}
