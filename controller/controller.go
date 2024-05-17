package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
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
var topCollections model.TopOpenSeaNFTCollections
var collectionSlug string
var mut sync.Mutex

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("can't load env file")
	}

	//For the websocket
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return r.Header.Get("Origin") == os.Getenv("TRUSTED_URL")
	}
}

/*
GetCollection returns info about a NFT Collection listed on Opensea.

Collection details include fees, traits, and links.
*/
func GetCollection(c *gin.Context) {

	collection = model.OpenSeaCollection{}
	collectionSlug = c.DefaultQuery("collection", "persona")
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
		log.Println("unable to unmarshal json")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, collection)
}

/*
GetNftStats returns stats about a single NFT Collection
*/
func GetNftStats(c *gin.Context) {

	collectionStats = model.OpenSeaCollectionStats{}
	collectionSlug = c.DefaultQuery("collection", "persona")
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
		log.Println("unable to unmarshal json")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, collectionStats)
}

/*
GetCollectionEvents returns the 50 most recent events for a NFT Collection
*/
func GetCollectionEvents(c *gin.Context) {
	event = model.OpenSeaCollectionEvent{}
	collectionSlug = c.DefaultQuery("collection", "persona")
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
		log.Println("unable to unmarshal json")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, event)
}

/*
GetTopNFTCollections returns the top NFT Collections on Opensea.com based on market cap.

Will return 50 NFTs by default, upper limit is 100
*/
func GetTopNFTCollections(c *gin.Context) {
	topCollections = model.TopOpenSeaNFTCollections{}
	count := c.DefaultQuery("limit", "50")
	url := "https://api.opensea.io/api/v2/collections?chain=ethereum&order_by=market_cap&limit=" + count

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

	if err := json.Unmarshal(body, &topCollections); err != nil {
		log.Println("unable to unmarshal json")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, topCollections)
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

/*
OpenSeaSocket provides a websocket connection to Opensea.com to subscribe to NTF events
in real-time.
*/
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
		Topic:   "collection:",
		Event:   "phx_join",
		Payload: struct{}{},
		Ref:     0,
	}

	go handleHeartbeat(openSeaConn)
	go updateSubscription(conn, openSeaConn, &subscribe)

	for {
		mt, message, err := openSeaConn.ReadMessage()
		if err != nil {
			log.Println("(OpenSea)read: ", err)
			break
		}
		mut.Lock()
		conn.WriteMessage(mt, message)
		mut.Unlock()
	}
}

/*
handleHeartbeat send a heartbeat message to Opensea.com websocket to keep the
connection alive in 30 second intervals.

Is safe to run concurrently.
*/
func handleHeartbeat(conn *websocket.Conn) {
	heartBeat := model.StreamHeartBeat{
		Topic:   "phoenix",
		Event:   "heartbeat",
		Payload: struct{}{},
		Ref:     0,
	}

	for {
		mut.Lock()
		conn.WriteJSON(heartBeat)
		mut.Unlock()
		time.Sleep(time.Millisecond * 30000)
		log.Println("heartbeat sent!")
	}
}

/*
updateSubscription allows the client to subscribe to NFT events.

Is safe to run concurrently.
*/
func updateSubscription(clientConn *websocket.Conn, openSeaConn *websocket.Conn, subscribe *model.StreamHeartBeat) {
	for {
		_, clientMsg, err := clientConn.ReadMessage()
		if err != nil {
			log.Println("(Client)read: ", err)
			break
		}
		subscribe.Topic = "collection:" + string(clientMsg)
		mut.Lock()
		openSeaConn.WriteJSON(subscribe)
		mut.Unlock()
	}
}
