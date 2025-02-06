package controllers

import (
	"OrderTracker/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLocation(context *gin.Context) {
	// This function will get a location from the database.
	// This function will be called from the routes.go file.

	// Get the location here.
	eventSlice, err := models.GetLocationQuery()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, eventSlice)
}

func PostLocation(context *gin.Context) {
	// This function will save a location to the database.
	// This function will be called from the routes.go file.

	// Save the location here.
	var location models.Location
	err := context.ShouldBindJSON(location)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = location.SaveLocationQuery()

	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "Location saved successfully"})

}

func PutLocation(context *gin.Context) {
	// This function will update a location in the database.
	// This function will be called from the routes.go file.

	// Update the location here.

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var location models.Location

	err = context.ShouldBindJSON(location)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = location.UpdateLocationQuery(id)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

}

func DeleteLocation(context *gin.Context) {
	// This function will delete a location from the database.
	// This function will be called from the routes.go file.

	// Delete the location here.
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	err = models.DeleteLocationQuery(id)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "Location deleted successfully"})

}

func GetLocationByID(context *gin.Context) {
	// This function will get a location by ID from the database.
	// This function will be called from the routes.go file.

	// Get the location by ID here.
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	location, err := models.GetLocationByIDQuery(id)
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, location)
}
