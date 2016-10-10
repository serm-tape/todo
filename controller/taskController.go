package controller

import (
	"github.com/labstack/echo"
	"github.com/serm-tape/todo/model"

	"net/http"
	"strconv"
)

const (
	bindFail = "Bind fail. Check your input format"
	createFail = "Create fail. Check your input value"
	updateFail = "Update fail. Check your input value"
)

func GetAllTask(c echo.Context) error{
	result := model.GetAllTask()
	return c.JSON(http.StatusOK, result)
}

func GetTaskById(c echo.Context) error{
	id := c.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Can not convert '"+id+"' to number")
	}
	return c.JSON(http.StatusOK, model.GetTaskById(uint(uid)))
}

func CreateTask(c echo.Context) error{
	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		return c.String(http.StatusBadRequest, bindFail)
	}

	if err := model.CreateTask(task); err != nil {
		return c.String(http.StatusBadRequest, createFail)
	}

	return c.JSON(http.StatusCreated, task)
}

func EditTask(c echo.Context) error{
	task := new(model.Task)

	id := c.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Can not convert '"+id+"' to number")
	}

	if err := c.Bind(task); err != nil {
		return c.String(http.StatusBadRequest, bindFail)
	}

	task.ID = uint(uid)
	if err := model.EditTask(task); err != nil {
		return c.String(http.StatusBadRequest, updateFail)
	}

	return c.JSON(http.StatusAccepted, task)
}

func DeleteTask(c echo.Context) error{
	id := c.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Can not convert '"+id+"' to number")
	}

	if err := model.DeleteTask(uint(uid)); err != nil {
		return c.String(http.StatusBadRequest, "Can not delete")
	}

	return c.String(http.StatusNoContent, "")
}