package controllers

import (
	"fmt"
	"os"
	"text/template"
)

func RouteFile() {
	templatePath := "F:\\Go_template_generator\\templateGen\\templates\\routeTemplate.txt"

	t, err := template.New("routeTemplate.txt").ParseFiles(templatePath)

	if err != nil {
		fmt.Println("Error parsing template:", err)
	}

	outputFile := "crud_app/routes/routes.go"
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
