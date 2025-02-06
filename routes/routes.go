package routes

import (
	"OrderTracker/controllers"
	"OrderTracker/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// This function will register the routes for the application.
	// This function will be called from the main.go file.

	// Register the routes here.

	//These are the routes protected by the middleware
	authenticstedRoutes := server.Group("/api/authenticated")
	authenticstedRoutes.Use(middleware.AuthMiddleware)

	//authenticstedRoutes.GET("/customer/orders", controllers.GetOrders)

	//These are the routes that are not protected by the middleware
	server.POST("/api/Location", controllers.PostLocation)
	server.GET("/api/Location", controllers.GetLocation)
	server.GET("/api/Location/:id", controllers.GetLocationByID)
	server.PUT("/api/Location/:id", controllers.PutLocation)
}
