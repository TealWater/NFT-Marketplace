package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/TealWater/NFT-Marketplace/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("can't load env file")
	}
}

func main() {
	fmt.Println("Hi mom!")

	router := gin.Default()

	router.GET("/", socket)
	router.GET("/opensea", openSeaSocket)

	router.Run(":8080")
}

func socket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade: ", err)
		return
	}
	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read: ", err)
			break
		}
		conn.WriteMessage(mt, message)

		val := 0
		for {
			msg := "message #" + strconv.Itoa(val)

			conn.WriteMessage(mt, []byte(msg))
			time.Sleep(time.Millisecond * 5000)
			val++
		}
	}
}

func openSeaSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade: ", err)
		return
	}
	defer conn.Close()

	openSeaDialer := websocket.Dialer{}
	url := "wss://stream.openseabeta.com/socket/websocket?token=" + os.Getenv("OPEN_SEA_KEY")
	openSeaConn, resp, err := openSeaDialer.DialContext(c, url, nil)
	if err != nil {
		log.Println("Open Sea Dialer: ", err)
		return
	}

	defer resp.Body.Close()
	defer openSeaConn.Close()

	subscribe := model.StreamHeartBeat{
		Topic:   "collection:BlockGames Dice",
		Event:   "phx_join",
		Payload: struct{}{},
		Ref:     0,
	}

	log.Println("here****---8")
	openSeaConn.WriteJSON(subscribe)
	go handleHeartBeat(openSeaConn)

	log.Println("here****---9")
	for {
		mt, message, err := openSeaConn.ReadMessage()
		if err != nil {
			log.Println("read: ", err)
			break
		}
		conn.WriteMessage(mt, message)
	}
}

func handleHeartBeat(conn *websocket.Conn) {
	log.Println("Made it -----")
	heartBeat := model.StreamHeartBeat{
		Topic:   "phoenix",
		Event:   "heartbeat",
		Payload: struct{}{},
		Ref:     0,
	}

	// ticker := time.NewTicker(time.Duration(30) * time.Second)
	//defer ticker.Stop()

	for {
		conn.WriteJSON(heartBeat)

		//ticker.Reset(30 * time.Second)
		time.Sleep(time.Millisecond * 30000)
		log.Println("heartbeat sent!")
	}
}
