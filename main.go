package main

import (
	"ShopifyChallenge/router"
)

func main() {
	//Get the router
	router := router.SetupRouter()

	//Start server on port 8080
	router.Run(":8080")
}
