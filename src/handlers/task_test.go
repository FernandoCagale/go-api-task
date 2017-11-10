package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/FernandoCagale/go-api-task/src/datastore"
	"github.com/FernandoCagale/go-api-task/src/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	db           *gorm.DB
	taskHandler  *TaskHandler
	ID           = 0
	taskJSONPost = `{"description":"Golang","type":"development"}`
	taskJSONPut  = `{"description":"Golang master","type":"development master"}`
)

func init() {
	db = getDb()
	taskHandler = NewTaskHandler()
}

func TestSaveTask(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/v1/tasks", strings.NewReader(taskJSONPost))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("db", db)

	task := models.Task{}

	// Assertions
	if assert.NoError(t, taskHandler.SaveTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		err := json.Unmarshal([]byte(rec.Body.String()), &task)

		assert.Nil(t, err)
		assert.NotNil(t, task.Id)
		assert.Equal(t, task.Description, "Golang")
		assert.Equal(t, task.Type, "development")
		ID = task.Id
	}
}

func TestGetAllTask(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/v1/tasks", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.Set("db", db)

	tasks := []models.Task{}

	// Assertions
	if assert.NoError(t, taskHandler.GetAllTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		err := json.Unmarshal([]byte(rec.Body.String()), &tasks)

		assert.Nil(t, err)
		assert.Equal(t, tasks[0].Id, ID)
		assert.Equal(t, tasks[0].Description, "Golang")
		assert.Equal(t, tasks[0].Type, "development")
	}
}

func TestGetTask(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/v1/tasks/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(ID))

	c.Set("db", db)

	task := models.Task{}

	// Assertions
	if assert.NoError(t, taskHandler.GetTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		err := json.Unmarshal([]byte(rec.Body.String()), &task)

		assert.Nil(t, err)
		assert.Equal(t, task.Id, ID)
		assert.Equal(t, task.Description, "Golang")
		assert.Equal(t, task.Type, "development")
	}
}

func TestUpdateTask(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/v1/tasks/:id", strings.NewReader(taskJSONPut))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(ID))

	c.Set("db", db)

	task := models.Task{}

	// Assertions
	if assert.NoError(t, taskHandler.UpdateTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		err := json.Unmarshal([]byte(rec.Body.String()), &task)

		assert.Nil(t, err)
		assert.NotNil(t, task.Id)
		assert.Equal(t, task.Description, "Golang master")
		assert.Equal(t, task.Type, "development master")
	}
}

func TestDeleteTask(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/v1/tasks/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(ID))

	c.Set("db", db)

	// Assertions
	if assert.NoError(t, taskHandler.DeleteTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		assert.Equal(t, `{"message":"Deleted"}`, rec.Body.String())
	}
}

func getDb() *gorm.DB {
	db, err := datastore.New("postgresql://postgres:postgres@localhost:5434/test_test?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.LogMode(false)

	return db
}
