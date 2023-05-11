package routes

import (
	"prakerja3/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.GET("/karyawans", controllers.GetKaryawanController)
	e.POST("/karyawans", controllers.InsertKaryawanController)
	return e
}
