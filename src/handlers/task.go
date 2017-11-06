package handlers

import (
	"net/http"

	"github.com/FernandoCagale/go-api-task/src/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func SaveTask(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	task := new(models.Task)

	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "BadRequest",
		})
	}

	if errors, valid := task.Validate(); !valid {
		return c.JSON(http.StatusBadRequest, errors)
	}

	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "InternalServerError",
		})
	}

	return c.JSON(http.StatusOK, task)
}

func GetAllTask(c echo.Context) error {
	tasks := []models.Task{}
	db := c.Get("db").(*gorm.DB)

	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "BadRequest",
		})
	}

	return c.JSON(http.StatusOK, tasks)
}

func GetTask(c echo.Context) error {
	id := c.Param("id")
	db := c.Get("db").(*gorm.DB)

	var task models.Task

	if err := db.Find(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "NotFound",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "InternalServerError",
		})
	}

	return c.JSON(http.StatusOK, task)
}

func UpdateTask(c echo.Context) error {
	id := c.Param("id")
	db := c.Get("db").(*gorm.DB)
	task := new(models.Task)

	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "BadRequest",
		})
	}

	if errors, valid := task.Validate(); !valid {
		return c.JSON(http.StatusBadRequest, errors)
	}

	if err := db.Find(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "NotFound",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "InternalServerError",
		})
	}

	if err := db.Model(&task).UpdateColumns(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "InternalServerError",
		})
	}

	return c.JSON(http.StatusOK, task)
}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	db := c.Get("db").(*gorm.DB)

	var task models.Task

	if err := db.Find(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "NotFound",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "InternalServerError",
		})
	}

	if err := db.Delete(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "InternalServerError",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Deleted",
	})
}
