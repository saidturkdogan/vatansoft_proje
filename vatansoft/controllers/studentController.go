package controllers

import (
	"net/http"
	"vatansoft/config"
	"vatansoft/models"

	"github.com/labstack/echo/v4"
)

func CreateStudent(c echo.Context) error {
	var student models.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	config.DB.Create(&student)
	return c.JSON(http.StatusOK, student)
}

func GetStudent(c echo.Context) error {
	id := c.Param("id")
	var student models.Student
	if result := config.DB.First(&student, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, student)
}

func UpdateStudent(c echo.Context) error {
	id := c.Param("id")
	var student models.Student
	if result := config.DB.First(&student, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	config.DB.Save(&student)
	return c.JSON(http.StatusOK, student)
}

func DeleteStudent(c echo.Context) error {
	id := c.Param("id")
	var student models.Student
	if result := config.DB.First(&student, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	config.DB.Delete(&student)
	return c.JSON(http.StatusOK, "Öğrenci silindi")
}
