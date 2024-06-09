package main

import (
	"fmt"
	"os"

	"github.com/TealWater/NFT-Marketplace/controller"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	fmt.Println("Hi mom!")

	router := gin.Default()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("TRUSTED_URL")},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Content-Type"},
	}))

	router.GET("/", controller.Socket)
	router.GET("/getNFT", controller.GetNftsByCollection)
	router.GET("/opensea", controller.OpenSeaSocket)
	router.GET("/getStats", controller.GetNftStats)
	router.GET("/getCollection", controller.GetCollection)
	router.GET("/getEvents", controller.GetCollectionEvents)
	router.GET("/getTopCollections", controller.GetTopNFTCollections)
	router.GET("/getGas", controller.GetEthGas)

	router.Run(":8080")
}
