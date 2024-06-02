package controllers

import (
	"net/http"
	"vatansoft/config"
	"vatansoft/models"

	"github.com/labstack/echo/v4"
)

func CreatePlan(c echo.Context) error {
	var plan models.Plan
	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// Çakışma kontrolü
	var existingPlan models.Plan
	if result := config.DB.Where("day = ? AND ((start_time <= ? AND end_time >= ?) OR (start_time <= ? AND end_time >= ?))", plan.Day, plan.StartTime, plan.StartTime, plan.EndTime, plan.EndTime).First(&existingPlan); result.RowsAffected > 0 {
		return c.JSON(http.StatusConflict, "Bu zaman aralığında zaten bir plan var")
	}
	config.DB.Create(&plan)
	return c.JSON(http.StatusOK, plan)
}

func GetPlan(c echo.Context) error {
	id := c.Param("id")
	var plan models.Plan
	if result := config.DB.First(&plan, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, plan)
}

func UpdatePlan(c echo.Context) error {
	id := c.Param("id")
	var plan models.Plan
	if result := config.DB.First(&plan, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	config.DB.Save(&plan)
	return c.JSON(http.StatusOK, plan)
}

func DeletePlan(c echo.Context) error {
	id := c.Param("id")
	var plan models.Plan
	if result := config.DB.First(&plan, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	config.DB.Delete(&plan)
	return c.JSON(http.StatusOK, "Plan silindi")
}
