package routes

import (
	"templateGen/controllers"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.POST("/insert-user", controllers.MainFunc)
}
