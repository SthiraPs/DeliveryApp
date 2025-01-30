package controllers

import (
	"github.com/SthiraPs/DeliveryApp/models"
	"github.com/SthiraPs/DeliveryApp/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateProductCategory godoc
// @Summary Create a new product category
// @Description Adds a new product category to the database
// @Tags productCategories
// @Accept json
// @Produce json
// @Param productCategory body models.ProductCategory true "ProductCategory data"
// @Success 201 {object} map[string]interface{} "ProductCategory created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/productCategory [post]
func CreateProductCategoryController(c *gin.Context) {
	var productCategory models.ProductCategory
	if err := c.ShouldBindJSON(&productCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to create the product category
	if err := services.CreateProductCategoryService(productCategory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product Category created successfully"})
}

// GetAllProductCategories godoc
// @Summary Get list of product categories
// @Description Returns all product categories from the database
// @Tags productCategories
// @Produce json
// @Success 200 {array} models.ProductCategory
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/productCategories [get]
func GetAllProductCategoriesController(c *gin.Context) {
	productCategories, err := services.GetAllProductCategoriesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productCategories)
}

// GetProductCategoryById godoc
// @Summary Get product category by ID
// @Description Returns a product category by its ID
// @Tags productCategories
// @Produce json
// @Param id path int true "Product Category ID"
// @Success 200 {object} models.ProductCategory
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/productCategory/{id} [get]
func GetProductCategoryByIdController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	productCategory, err := services.GetProductCategoryByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productCategory)
}

// UpdateProductCategoryById godoc
// @Summary Update a product category
// @Description Updates product category information by ID
// @Tags productCategories
// @Accept json
// @Produce json
// @Param id path int true "Product Category ID"
// @Param productCategory body models.ProductCategory true "Updated product category data"
// @Success 200 {object} models.ProductCategory "Product category updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 404 {object} map[string]interface{} "Product category not found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/productCategory/{id} [put]
func UpdateProductCategoryByIdController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedProductCategory models.ProductCategory
	if err := c.ShouldBindJSON(&updatedProductCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productCategory, err := services.UpdateProductCategoryByIdService(id, updatedProductCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, productCategory)
}

// DeleteProductCategoryById godoc
// @Summary Delete a product category
// @Description Deletes a product category by ID
// @Tags productCategories
// @Param id path int true "Product Category ID"
// @Success 200 {object} map[string]interface{} "Product category deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 404 {object} map[string]interface{} "Product category not found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/productCategory/{id} [delete]
func DeleteProductCategoryByIdController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	productCategory, err := services.DeleteProductCategoryByIdService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product Category deleted successfully", "productCategory": productCategory})
}
