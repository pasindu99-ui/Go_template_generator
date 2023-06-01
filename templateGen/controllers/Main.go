package controllers

import (
	"log"
	"os"

	"github.com/labstack/echo"
)

func MainFunc(c echo.Context) error {
	if err := os.MkdirAll("crud_app/configs", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("crud_app/controllers", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("crud_app/models", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("crud_app/routes", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	mongoURL := c.FormValue("mongoURL")
	databaseName := c.FormValue("databaseName")
	collectionName := c.FormValue("collectionName")
	EnvFile()
	ServerFile()
	RouteFile()
	SetupFile()
	InsertUser(mongoURL, databaseName, collectionName)
	UpdateUser(mongoURL, databaseName, collectionName)
	DeleteUser(mongoURL, databaseName, collectionName)
	ViewUser(mongoURL, databaseName, collectionName)
	return c.String(200, "Main function executed successfully")

}
