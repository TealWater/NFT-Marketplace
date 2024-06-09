package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/TealWater/NFT-Marketplace/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum/ethclient"
)

var upgrader = websocket.Upgrader{}
var collectionStats model.OpenSeaCollectionStats
var collection model.OpenSeaCollection
var event model.OpenSeaCollectionEvent
var topCollections model.TopOpenSeaNFTCollections
var collectionSlug string
var mut sync.Mutex
var ethConn *ethclient.Client

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("can't load env file")
	}

	//For the websocket
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return r.Header.Get("Origin") == os.Getenv("TRUSTED_URL")
	}

	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/" + os.Getenv("INFRA_IO_KEY"))
	if err != nil {
		log.Println("unable to connect to ethereum node: ", err)
	}
	ethConn = conn
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

func GetNftsByCollection(c *gin.Context) {
	nftListing := &model.OpenSeaNFTListing{}
	nfts := &model.OpenSeaNFT{}
	collection := c.DefaultQuery("collection", "persona")
	next := ""
	// isDone := false

	//for loop will be used for pagination
	// for !isDone {
	url := "https://api.opensea.io/api/v2/listings/collection/" + collection + "/all?limit=100&next=" + next

	currentNFTListingRequest := &model.OpenSeaNFTListing{}

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

	if err := json.Unmarshal(body, currentNFTListingRequest); err != nil {
		log.Println("unable to unmarshal json")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// log.Println(nft.Next)

	nftListing.Listings = append(nftListing.Listings, currentNFTListingRequest.Listings...)

	next = currentNFTListingRequest.Next
	// if len(next) == 0 {
	// 	isDone = true
	// }

	// }

	for idx := range nftListing.Listings {
		singleNFT, err := getSingleOpenSeaNFT(c, *nftListing, idx)
		if err != nil {
			log.Println("There is an issue with getting a single nft")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		nfts.Nfts = append(nfts.Nfts, singleNFT.Nfts...)
	}
	c.JSON(http.StatusOK, nfts)
}

func getSingleOpenSeaNFT(c *gin.Context, nft model.OpenSeaNFTListing, idx int) (model.OpenSeaNFT, error) {
	chain := "ethereum"
	address := nft.Listings[idx].ProtocolData.Parameters.Offer[0].Token
	identifier := nft.Listings[idx].ProtocolData.Parameters.Offer[0].IdentifierOrCriteria
	singleOpenSeaNFT := &model.SingleOpenSeaNFT{}
	builtNFT := model.OpenSeaNFT{}

	url := "https://api.opensea.io/api/v2/chain/" + chain + "/contract/" + address + "/nfts/" + identifier
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return builtNFT, err
	}

	req.Header = http.Header{
		"accept":    {"application/json"},
		"x-api-key": {os.Getenv("OPEN_SEA_KEY")},
	}

	res, err := client.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return builtNFT, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return builtNFT, err
	}

	if err := json.Unmarshal(body, singleOpenSeaNFT); err != nil {
		log.Println("unable to unmarshal json")
		log.Println("inner method! getOpenSeaNFT")
		c.AbortWithError(http.StatusInternalServerError, err)
		return builtNFT, err
	}

	//build the nft
	//can append an nft struct to an empty slice
	builtNFT.Nfts = make([]struct {
		Identifier    string  "json:\"identifier\""
		Collection    string  "json:\"collection\""
		Contract      string  "json:\"contract\""
		TokenStandard string  "json:\"token_standard\""
		Name          string  "json:\"name\""
		Description   string  "json:\"description\""
		ImageURL      string  "json:\"image_url\""
		MetadataURL   string  "json:\"metadata_url\""
		OpenseaURL    string  "json:\"opensea_url\""
		UpdatedAt     string  "json:\"updated_at\""
		IsDisabled    bool    "json:\"is_disabled\""
		IsNsfw        bool    "json:\"is_nsfw\""
		Price         float64 "json:\"price\""
		Currency      string  "json:\"currency\""
	}, 1)
	builtNFT.Nfts[0].Identifier = singleOpenSeaNFT.Nft.Identifier
	builtNFT.Nfts[0].Collection = singleOpenSeaNFT.Nft.Collection
	builtNFT.Nfts[0].Contract = singleOpenSeaNFT.Nft.Contract
	builtNFT.Nfts[0].TokenStandard = singleOpenSeaNFT.Nft.TokenStandard
	builtNFT.Nfts[0].Name = singleOpenSeaNFT.Nft.Name
	builtNFT.Nfts[0].Description = singleOpenSeaNFT.Nft.Description
	builtNFT.Nfts[0].ImageURL = singleOpenSeaNFT.Nft.ImageURL
	builtNFT.Nfts[0].MetadataURL = singleOpenSeaNFT.Nft.MetadataURL
	builtNFT.Nfts[0].OpenseaURL = singleOpenSeaNFT.Nft.OpenseaURL
	builtNFT.Nfts[0].UpdatedAt = singleOpenSeaNFT.Nft.UpdatedAt
	builtNFT.Nfts[0].IsDisabled = singleOpenSeaNFT.Nft.IsDisabled
	builtNFT.Nfts[0].IsNsfw = singleOpenSeaNFT.Nft.IsNsfw

	stringPrice := nft.Listings[idx].Price.Current.Value
	decimal := nft.Listings[idx].Price.Current.Decimals
	decimalPlacement := len(stringPrice) - decimal
	//len(str) - decimal --> decimal placement

	// log.Println("price: ", stringPrice, "\ndecimals: ", decimal, "\ndecimal placement: ", decimalPlacement, "\nstring len: ", len(stringPrice))

	var buffer bytes.Buffer
	if decimalPlacement < 0 {
		buffer.WriteString("0.")
		for decimalPlacement < 0 {
			buffer.WriteString("0")
			decimalPlacement++
		}
		buffer.WriteString(stringPrice)
	} else {
		buffer.WriteString(stringPrice[:decimalPlacement])
		buffer.WriteString(".")
		buffer.WriteString(stringPrice[decimalPlacement:])
	}

	builtNFT.Nfts[0].Price, _ = strconv.ParseFloat(buffer.String(), 32)
	builtNFT.Nfts[0].Currency = nft.Listings[idx].Price.Current.Currency
	return builtNFT, nil
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

/*
GetEthGas - returns the minimun amount of gas required to complete a transaction on
the block, actual gas requre may vary at the time of the transaction

TODO: add safeExit() func to return gracefully if there is an error
*/
func GetEthGas(c *gin.Context) {
	c.Request.Header.Add("Content-Type:", "application/json")
	val, err := ethConn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("unable to fetch a gas price: ", err)
	}
	ethGasInWEI := new(big.Float).SetInt(val)
	oneBilliion := big.NewFloat(1000000000)
	ethGasInGwei := new(big.Float).Quo(ethGasInWEI, oneBilliion)

	c.JSON(http.StatusAccepted, fmt.Sprintf("%.2f gWEI", ethGasInGwei))
}
