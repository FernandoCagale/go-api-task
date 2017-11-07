package main

import (
	"github.com/FernandoCagale/go-api-task/src/checker"
	"github.com/FernandoCagale/go-api-task/src/datastore"
	"github.com/FernandoCagale/go-api-task/src/handlers"
	"github.com/FernandoCagale/go-api-task/src/lib"

	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db, err := datastore.ConnectDB()
	failOnError(err, "Failed to init dababase connection!")
	defer db.Close()

	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(lib.BindDb(db))

	checkers := map[string]checker.Checker{
		"api":      checker.NewApi(),
		"postgres": checker.NewPostgres(datastore.Connection()),
	}

	healthzHandler := handlers.NewHealthzHandler(checkers)
	app.GET("/health", healthzHandler.HealthzIndex)

	group := app.Group("/v1")

	tasksHandler := handlers.NewTaskHandler()

	group.GET("/tasks", tasksHandler.GetAllTask)
	group.POST("/tasks", tasksHandler.SaveTask)
	group.GET("/tasks/:id", tasksHandler.GetTask)
	group.PUT("/tasks/:id", tasksHandler.UpdateTask)
	group.DELETE("/tasks/:id", tasksHandler.DeleteTask)

	app.Logger.Fatal(app.Start(":3000"))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Error(err)
		log.Info(msg)
	}
}
