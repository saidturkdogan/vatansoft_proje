package routes

import (
	"vatansoft/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/plans", controllers.CreatePlan)
	e.GET("/plans/:id", controllers.GetPlan)
	e.PUT("/plans/:id", controllers.UpdatePlan)
	e.DELETE("/plans/:id", controllers.DeletePlan)
	e.POST("/students", controllers.CreateStudent)
	e.GET("/students/:id", controllers.GetStudent)
	e.PUT("/students/:id", controllers.UpdateStudent)
	e.DELETE("/students/:id", controllers.DeleteStudent)
}
