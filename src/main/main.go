package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Task struct {
	Description string `json:"description"`
	Type string `json:"type"`
}

func (t *Task) validate() (errors map[string]string, ok bool) {
	errors = make(map[string]string)

	if (t.Description == "") {
		errors["description"] = "field description invalid"
	}

	if (t.Type == "") {
		errors["type"] = "field type invalid"
	}

	ok = len(errors) == 0

	return
}

func saveTask(c echo.Context) error {
	u := new(Task)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Bad Request",
		})
	}

	if errors, valid := u.validate(); !valid {
		return c.JSON(http.StatusInternalServerError, errors)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"POST":     "saveTask",
		"description": u.Description,
		"type": u.Type,
	})
}

func getTask(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"GET": "getTask",
		"id":  id,
	})
}

func updateTask(c echo.Context) error {
	id := c.Param("id")
	u := new(Task)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Bad Request",
		})
	}

	if errors, valid := u.validate(); !valid {
		return c.JSON(http.StatusInternalServerError, errors)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"PUT":      "putTask",
		"id":       id,
		"description": u.Description,
		"type": u.Type,
	})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]string{
		"DELETE": "deleteTask",
		"id":     id,
	})
}

func main() {
	e := echo.New()

	e.POST("/tasks", saveTask)
	e.GET("/tasks/:id", getTask)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Logger.Fatal(e.Start(":8000"))
}
