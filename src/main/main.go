package main

import (
	"net/http"

	shared "github.com/FernandoCagale/go-api-shared/src/validation"
	"github.com/labstack/echo"
)

type Task struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

func (t *Task) validate() (errors map[string][]shared.Validation, ok bool) {
	errors = make(map[string][]shared.Validation)

	if t.Description == "" {
		errors["description"] = append(errors["description"], shared.Validation{
			Type:    "required",
			Message: "field is required",
		})
	}

	if len(t.Description) > 40 {
		errors["description"] = append(errors["description"], shared.Validation{
			Type:    "lenght-max",
			Message: "field lenght max 40",
		})
	}

	if len(t.Description) < 5 {
		errors["description"] = append(errors["description"], shared.Validation{
			Type:    "lenght-min",
			Message: "field lenght min 5",
		})
	}

	if t.Type == "" {
		errors["type"] = append(errors["type"], shared.Validation{
			Type:    "required",
			Message: "field is required",
		})
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
		"POST":        "saveTask",
		"description": u.Description,
		"type":        u.Type,
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
		"PUT":         "updateTask",
		"id":          id,
		"description": u.Description,
		"type":        u.Type,
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
