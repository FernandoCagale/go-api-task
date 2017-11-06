package routers

import (
	"github.com/FernandoCagale/go-api-task/src/handlers"
	"github.com/labstack/echo"
)

var App *echo.Echo

func init() {
	App = echo.New()

	group := App.Group("/v1")

	group.GET("/tasks", handlers.GetAllTask)
	group.POST("/tasks", handlers.SaveTask)
	group.GET("/tasks/:id", handlers.GetTask)
	group.PUT("/tasks/:id", handlers.UpdateTask)
	group.DELETE("/tasks/:id", handlers.DeleteTask)
}
