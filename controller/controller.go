package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/TealWater/NFT-Marketplace/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var upgrader = websocket.Upgrader{}
var collectionStats model.OpenSeaCollectionStats
var collection model.OpenSeaCollection
var event model.OpenSeaCollectionEvent

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("can't load env file")
	}
}

func GetCollection(c *gin.Context) {

	collection = model.OpenSeaCollection{}
	collectionSlug := "persona"
	url := "https://api.opensea.io/api/v2/collections/" + collectionSlug

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req.Header = http.Header{
		"accept":    {"application/json"},
		"x-api-key": {os.Getenv("OPEN_SEA_KEY")},
	}

	res, err := client.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := json.Unmarshal(body, &collection); err != nil {
		log.Println("unable to unmarshal json\n")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, collection)
}

func GetNftStats(c *gin.Context) {

	collectionStats = model.OpenSeaCollectionStats{}
	collectionSlug := "persona"
	url := "https://api.opensea.io/api/v2/collections/" + collectionSlug + "/stats"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req.Header = http.Header{
		"accept":    {"application/json"},
		"x-api-key": {os.Getenv("OPEN_SEA_KEY")},
	}

	res, err := client.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := json.Unmarshal(body, &collectionStats); err != nil {
		log.Println("unable to unmarshal json\n")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, collectionStats)
}

func GetCollectionEvents(c *gin.Context) {
	event = model.OpenSeaCollectionEvent{}
	collectionSlug := "persona"
	url := "https://api.opensea.io/api/v2/events/collection/" + collectionSlug

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req.Header = http.Header{
		"accept":    {"application/json"},
		"x-api-key": {os.Getenv("OPEN_SEA_KEY")},
	}

	res, err := client.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := json.Unmarshal(body, &event); err != nil {
		log.Println("unable to unmarshal json\n")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, event)
}

func Socket(c *gin.Context) {
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

func OpenSeaSocket(c *gin.Context) {
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
		Topic:   "collection:persona",
		Event:   "phx_join",
		Payload: struct{}{},
		Ref:     0,
	}

	openSeaConn.WriteJSON(subscribe)
	go HandleHeartBeat(openSeaConn)
	go UpdateSubscription(conn, openSeaConn, &subscribe)

	for {
		mt, message, err := openSeaConn.ReadMessage()
		if err != nil {
			log.Println("(OpenSea)read: ", err)
			break
		}
		conn.WriteMessage(mt, message)
	}
}

func HandleHeartBeat(conn *websocket.Conn) {
	heartBeat := model.StreamHeartBeat{
		Topic:   "phoenix",
		Event:   "heartbeat",
		Payload: struct{}{},
		Ref:     0,
	}

	for {
		conn.WriteJSON(heartBeat)
		time.Sleep(time.Millisecond * 30000)
		log.Println("heartbeat sent!")
	}
}

func UpdateSubscription(clientConn *websocket.Conn, openSeaConn *websocket.Conn, subscribe *model.StreamHeartBeat) {
	for {
		_, clientMsg, err := clientConn.ReadMessage()
		if err != nil {
			log.Println("(Client)read: ", err)
			break
		}
		subscribe.Topic = "collection:" + string(clientMsg)
		openSeaConn.WriteJSON(subscribe)
	}
}
