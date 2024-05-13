package main

import (
	"fmt"

	"github.com/TealWater/NFT-Marketplace/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi mom!")

	router := gin.Default()

	router.GET("/", controller.Socket)
	router.GET("/opensea", controller.OpenSeaSocket)
	router.GET("/getStats", controller.GetNftStats)
	router.GET("/getCollection", controller.GetCollection)
	router.GET("/getEvents", controller.GetCollectionEvents)

	router.Run(":8080")
}
