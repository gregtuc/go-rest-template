package router

import (
	"ShopifyChallenge/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//Default crud requirements
	router.POST("/items/create", controllers.CreateItem)
	router.GET("/items/read", controllers.ReadItems)
	router.PATCH("/items/update", controllers.UpdateItem)
	router.DELETE("/items/delete/:id", controllers.DeleteItem)

	//Additional requirement: "Ability to create “shipments” and assign inventory to the shipment, and adjust inventory appropriately"
	router.POST("/shipments/create", controllers.CreateShipment)
	router.GET("/shipments/read", controllers.ReadShipments)

	return router
}
