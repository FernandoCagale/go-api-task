package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	taskJsonPOST    = `{"POST":"saveTask","description":"Golang","type":"development"}`
	taskJsonInvalid = `{"description":[{"type":"required","message":"field is required"},{"type":"lenght-min","message":"field lenght min 5"}],"type":[{"type":"required","message":"field is required"}]}`
	taskJsonPUT     = `{"PUT":"updateTask","description":"Golang","id":"1","type":"development"}`
	taskJsonGET     = `{"GET":"getTask","id":"1"}`
	taskJsonDELETE  = `{"DELETE":"deleteTask","id":"1"}`
)

func TestSaveTask(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/tasks", strings.NewReader(`{"description":"Golang","type":"development"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, saveTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, taskJsonPOST, rec.Body.String())
	}
}

func TestSaveTaskValidate(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/tasks", strings.NewReader(`{"description":"","type":""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, saveTask(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, taskJsonInvalid, rec.Body.String())
	}
}

func TestUpdateTask(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/tasks/:id", strings.NewReader(`{"description":"Golang","type":"development"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, updateTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, taskJsonPUT, rec.Body.String())
	}
}

func TestUpdateTaskValidate(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/tasks/:id", strings.NewReader(`{"description":"","type":""}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, updateTask(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, taskJsonInvalid, rec.Body.String())
	}
}
func TestDeleteTask(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/tasks/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, deleteTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, taskJsonDELETE, rec.Body.String())
	}
}

func TestGetTask(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/tasks/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, getTask(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, taskJsonGET, rec.Body.String())
	}
}
