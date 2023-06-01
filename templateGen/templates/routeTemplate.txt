package routes

import (
	"Crud_app/controllers"

	"github.com/labstack/echo"
)

func UserRoute(e *echo.Echo) {
	e.POST("/create-user", controllers.InsertUser)
	e.POST("/update-user", controllers.UpdateUser)
	e.DELETE("/delete-user", controllers.DeleteUser)
	e.GET("/view-user", controllers.ViewUser)

}
