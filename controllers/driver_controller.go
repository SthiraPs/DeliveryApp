package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/SthiraPs/DeliveryApp/models"
	"github.com/SthiraPs/DeliveryApp/services"
	"github.com/gin-gonic/gin"
)

// CreateDriver godoc
// @Summary Create a new driver
// @Description Adds a new driver to the database
// @Tags drivers
// @Accept json
// @Produce json
// @Param driver body models.Driver true "Driver data"
// @Success 201 {object} map[string]interface{} "Driver created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/driver [post]
func CreateDriver(c *gin.Context) {
	var driver models.Driver
	if err := c.ShouldBindJSON(&driver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driver.CreatedAt = time.Now()
	driver.UpdatedAt = time.Now()

	if err := services.CreateDriverService(driver); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Driver created successfully"})
}

// GetAllDrivers godoc
// @Summary Get list of drivers
// @Description Returns all drivers from the database
// @Tags drivers
// @Produce json
// @Success 200 {array} models.Driver
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/drivers [get]
func GetAllDrivers(c *gin.Context) {
	drivers, err := services.GetAllDriversService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drivers)
}

// GetDriverById godoc
// @Summary Get driver by ID
// @Description Returns a driver by their ID
// @Tags drivers
// @Produce json
// @Param id path int true "Driver ID"
// @Success 200 {object} models.Driver
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/driver/{id} [get]
func GetDriverById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	driver, err := services.GetDriverByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, driver)
}

// UpdateDriverById godoc
// @Summary Update a driver
// @Description Updates driver information by ID
// @Tags drivers
// @Accept json
// @Produce json
// @Param id path int true "Driver ID"
// @Param driver body models.Driver true "Updated driver data"
// @Success 200 {object} models.Driver "Driver updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 404 {object} map[string]interface{} "Driver not found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/driver/{id} [put]
func UpdateDriverById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedDriver models.Driver
	if err := c.ShouldBindJSON(&updatedDriver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driver, err := services.UpdateDriverByIdServer(id, updatedDriver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, driver)
}

// DeleteDriverById godoc
// @Summary Delete a driver
// @Description Deletes a driver by ID
// @Tags drivers
// @Param id path int true "Driver ID"
// @Success 200 {object} map[string]interface{} "Driver deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 404 {object} map[string]interface{} "Driver not found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/driver/{id} [delete]
func DeleteDriverById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	driver, err := services.DeleteDriverByIdServer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Driver deleted successfully", "driver": driver})
}
