package main

import (
	"time"

	"github.com/FernandoCagale/go-api-task/src/checker"
	"github.com/FernandoCagale/go-api-task/src/config"
	"github.com/FernandoCagale/go-api-task/src/datastore"
	"github.com/FernandoCagale/go-api-task/src/handlers"
	"github.com/FernandoCagale/go-api-task/src/lib"
	"github.com/jinzhu/gorm"

	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	var db *gorm.DB
	env := config.LoadEnv()
	app := echo.New()

	go bindDatastore(app, db, env.DatastoreURL)

	app.Use(middleware.Logger())

	defer db.Close()

	checkers := map[string]checker.Checker{
		"api":      checker.NewApi(),
		"postgres": checker.NewPostgres(env.DatastoreURL),
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

	app.Logger.Fatal(app.Start(":" + env.Port))
}

func bindDatastore(app *echo.Echo, db *gorm.DB, url string) {
	for {
		db, err := datastore.New(url)
		failOnError(err, "Failed to init dababase connection!")
		if err == nil {
			app.Use(lib.BindDb(db))
			break
		}
		time.Sleep(time.Second * 5)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Info(msg)
	}
}
