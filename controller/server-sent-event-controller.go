package controller

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

var Stream *Event
var ethConn *ethclient.Client

func init() {
	Stream = NewServer()

	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/" + os.Getenv("INFRA_IO_KEY"))
	if err != nil {
		log.Println("unable to connect to ethereum node: ", err)
	}
	ethConn = conn

	/*
		to allow one source of eth gas updating, prevents having n routines for n client connections
	*/
	go handleGasUpdate()
}

type Event struct {
	// Events are pushed to this channel by the main events-gathering routine
	Message chan string

	// New client connections
	NewClients chan chan string

	// Closed client connections
	ClosedClients chan chan string

	// Total client connections
	TotalClients map[chan string]bool
}

// New event messages are broadcast to all registered client connection channels
type ClientChan chan string

func (stream *Event) ServeHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Initialize client channel
		clientChan := make(ClientChan)

		//Send new connection to event server
		stream.NewClients <- clientChan

		defer func() {
			// Send closed connection to event server
			stream.ClosedClients <- clientChan
		}()

		c.Set("clientChan", clientChan)
		c.Writer.Flush()
		c.Next()
	}
}

/*
HandleMiddleware - sets all of the headers needed for creating a server-side-event
or a server-send-event server
*/
func HandleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Writer.Flush()
		c.Next()
	}
}

func (stream *Event) listen() {
	for {
		select {
		// Add new available client
		case client := <-stream.NewClients:
			stream.TotalClients[client] = true
			log.Printf("Client added. %d registered clients", len(stream.TotalClients))

		// Remove closed client
		case client := <-stream.ClosedClients:
			delete(stream.TotalClients, client)
			close(client)
			log.Printf("Removed client, %d registered clients", len(stream.TotalClients))

		// Broadcast message to client
		case eventMsg := <-stream.Message:
			for clientMessageChan := range stream.TotalClients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

/*
NewServer - returns a pointer to an event struct and allows the event server to listen
for incomming client connections
*/
func NewServer() (event *Event) {
	event = &Event{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}
	go event.listen()
	return
}

/*
getEthGas - returns the minimun amount of gas required to complete a transaction on
the block, actual gas requre may vary at the time of the transaction

TODO: add safeExit() func to return gracefully if there is an error
*/
func getEthGas() string {
	val, err := ethConn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("unable to fetch a gas price: ", err)
	}
	ethGasInWEI := new(big.Float).SetInt(val)
	oneBilliion := big.NewFloat(1000000000)
	ethGasInGwei := new(big.Float).Quo(ethGasInWEI, oneBilliion)

	// log.Printf("%.2f gWEI \n\n", ethGasInGwei)
	return fmt.Sprintf("%.2f gWEI \n\n", ethGasInGwei)
}

func handleGasUpdate() {
	for {
		Stream.Message <- getEthGas()
		time.Sleep(time.Second * 30)
	}
}

/*
StreamGasPrice - sends a server event to update the client on the current price of Eth gas
*/
func StreamGasPrice(c *gin.Context) {
	v, ok := c.Get("clientChan")
	if !ok {
		return
	}
	clientChan, ok := v.(ClientChan)
	if !ok {
		return
	}
	c.Stream(func(w io.Writer) bool {
		// Stream message to client from message channel
		if msg, ok := <-clientChan; ok {
			c.SSEvent("message", msg)
			c.Writer.Flush()
			return true
		}
		return false
	})
}
