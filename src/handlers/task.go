package handlers

import (
	"net/http"

	"github.com/FernandoCagale/go-api-task/src/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type TaskHandler struct {
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) SaveTask(c echo.Context) error {
	task := new(models.Task)

	db, valid := getConnection(c)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "InternalServerError",
		})
	}

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

func (h *TaskHandler) GetAllTask(c echo.Context) error {
	tasks := []models.Task{}

	db, valid := getConnection(c)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "InternalServerError",
		})
	}

	query := models.NewQuery(c)

	where := make(map[string]interface{})

	if !query.IsEmpty() {
		where["description"] = query.Search
	}

	if err := db.Limit(query.GetLimit()).Offset(query.Offset).Where(where).Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "InternalServerError",
		})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTask(c echo.Context) error {
	id := c.Param("id")
	var task models.Task

	db, valid := getConnection(c)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "InternalServerError",
		})
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

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	task := new(models.Task)

	db, valid := getConnection(c)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "InternalServerError",
		})
	}

	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "BadRequest",
		})
	}

	if errors, valid := task.Validate(); !valid {
		return c.JSON(http.StatusBadRequest, errors)
	}

	if err := db.Find(&models.Task{}, id).Error; err != nil {
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

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	var task models.Task

	db, valid := getConnection(c)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "InternalServerError",
		})
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

	if err := db.Delete(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "InternalServerError",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Deleted",
	})
}

func getConnection(c echo.Context) (*gorm.DB, bool) {
	db := c.Get("db")
	if db != nil {
		return db.(*gorm.DB), true
	}
	return nil, false
}
