package routes

import (
	"templateGen/controllers"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.POST("/insert-user", controllers.InsertUser)
	e.POST("/update-user", controllers.UpdateUser)
	e.POST("/view-user", controllers.ViewUser)
	e.POST("/delete-user", controllers.DeleteUser)
}
