package controllers

import (
	"OrderTracker/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(context *gin.Context) {
	// This function will get a product from the database.
	// This function will be called from the routes.go file.

	// Get the product here.
	productSlice, err := models.GetProductQuery()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, productSlice)
}

func PostProduct(context *gin.Context) {
	// This function will save a product to the database.
	// This function will be called from the routes.go file.

	// Save the product here.
	var product models.Product
	err := context.ShouldBindJSON(&product)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = product.SaveProductQuery()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "Product saved successfully"})

}
