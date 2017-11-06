package main

import (
	"github.com/FernandoCagale/go-api-task/src/datastore"
	"github.com/FernandoCagale/go-api-task/src/lib"
	"github.com/FernandoCagale/go-api-task/src/routers"
	"github.com/labstack/echo/middleware"
)

func main() {
	db := datastore.ConnectDB()

	defer db.Close()

	e := routers.App
	e.Use(middleware.Logger())
	e.Use(lib.BindDb(db))

	e.Logger.Fatal(e.Start(":3000"))
}
