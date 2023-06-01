package controllers

import (
	"fmt"
	"os"
	"text/template"
)

func SetupFile() {
	templatePath := "F:\\Go_template_generator\\templateGen\\templates\\setupTemplate.txt"

	t, err := template.New("setupTemplate.txt").ParseFiles(templatePath)

	if err != nil {
		fmt.Println("Error parsing template:", err)
	}

	outputFile := "crud_app/configs/setup.go"
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
