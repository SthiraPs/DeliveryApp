package controllers

import (
	"github.com/SthiraPs/DeliveryApp/models"
	"github.com/SthiraPs/DeliveryApp/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateOrder godoc
// @Summary Create a new order
// @Description Adds a new order to the database, including associated order items
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order data"
// @Success 201 {object} map[string]interface{} "Order created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/order [post]
func CreateOrderController(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to create the order
	if err := services.CreateOrderService(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}

// GetAllOrders godoc
// @Summary Get list of orders
// @Description Returns all orders from the database
// @Tags orders
// @Produce json
// @Success 200 {array} models.Order
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/orders [get]
func GetAllOrdersController(c *gin.Context) {
	orders, err := services.GetAllOrdersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetOrderById godoc
// @Summary Get order by ID
// @Description Returns an order by its ID
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/order/{id} [get]
func GetOrderByIdController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	order, err := services.GetOrderByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

// UpdateOrderById godoc
// @Summary Update an order
// @Description Updates order information by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body models.Order true "Updated order data"
// @Success 200 {object} models.Order "Order updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 404 {object} map[string]interface{} "Order not found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/order/{id} [put]
func UpdateOrderByIdController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedOrder models.Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := services.UpdateOrderByIdService(id, updatedOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// DeleteOrderById godoc
// @Summary Delete an order
// @Description Deletes an order by ID
// @Tags orders
// @Param id path int true "Order ID"
// @Success 200 {object} map[string]interface{} "Order deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 404 {object} map[string]interface{} "Order not found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/order/{id} [delete]
func DeleteOrderByIdController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	order, err := services.DeleteOrderByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully", "order": order})
}
