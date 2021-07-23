package lanyard

import (
	"log"
	"strings"
	"time"

	"github.com/sacOO7/gowebsocket"
)

const (
	PING_PERIOD = 30 * time.Second
)

type WSClient struct {
	socket gowebsocket.Socket
	ticker *time.Ticker
}

func (client WSClient) Destroy() {
	if client.ticker != nil {
		client.ticker.Stop()
	}
	client.socket.Close()
}

func (client WSClient) ping() {
	client.ticker = time.NewTicker(PING_PERIOD)
	defer client.ticker.Stop()

	for ; ; <-client.ticker.C {
		client.socket.SendText("{\"op\":3}")
	}
}

func clientFactory(subscription string, subscription_data string, onMessage func(client WSClient, message string)) *WSClient {
	client := WSClient{
		socket: gowebsocket.New(WS_URL),
	}

	client.socket.OnConnected = func(socket gowebsocket.Socket) {
		client.socket.SendText("{\"op\":2,\"d\":{\"" + subscription + "\":" + subscription_data + "}}") 
		go client.ping()
	}

	client.socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Println("An error occured while connecting to Lanyard websocket server", err)
		client.Destroy()
	}

	client.socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		if strings.Contains(message, "heartbeat_interval") {
			return
		}
		
		onMessage(client, message)
	}

	client.socket.Connect()

	return &client
}
