package controllers

import (
	"bytes"
	"net/http"
	"text/template"

	"github.com/labstack/echo"
)

type Details struct {
	MongoURL       string
	DatabaseName   string
	CollectionName string
}

func InsertUser(c echo.Context) error {
	mongoURL := c.FormValue("mongoURL")
	databaseName := c.FormValue("databaseName")
	collectionName := c.FormValue("collectionName")

	details := Details{mongoURL, databaseName, collectionName}
	templatePath := "G:\\Go_template_generator\\templateGen\\templates\\insertTemplate.txt"

	t, err := template.New("insertTemplate.txt").ParseFiles(templatePath)

	if err != nil {
		return err
	}
	var buffer bytes.Buffer

	err = t.Execute(&buffer, details)
	if err != nil {
		return err
	}
	generatedCode := buffer.String()

	return c.String(http.StatusOK, generatedCode)
}
