package lanyard

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/sacOO7/gowebsocket"
)

const (
	WS_URL      = "wss://api.lanyard.rest/socket"
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

func ListenUser(userId string, presenceUpdate func(data *LanyardData)) WSClient {
	client := WSClient{
		socket: gowebsocket.New(WS_URL),
	}

	client.socket.OnConnected = func(socket gowebsocket.Socket) {
		client.socket.SendText("{\"op\":2,\"d\":{\"subscribe_to_id\":\"" + userId + "\"}}")
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

		var data LanyardWSResponse
		err := json.Unmarshal([]byte(message), &data)
		if err != nil {
			client.Destroy()
			return
		}

		presenceUpdate(data.D)
	}

	client.socket.Connect()

	return client
}

func ListenMultipleUsers(userIds []string, presenceUpdate func(data []*LanyardData)) WSClient {
	client := WSClient{
		socket: gowebsocket.New(WS_URL),
	}

	var formattedIds []string

	for _, id := range userIds {
		formattedIds = append(formattedIds, "\""+id+"\"")
	}

	client.socket.OnConnected = func(socket gowebsocket.Socket) {
		client.socket.SendText("{\"op\":2,\"d\":{\"subscribe_to_ids\":[" + strings.Join(formattedIds, ",") + "]}}")
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

		var data map[string]json.RawMessage

		err := json.Unmarshal([]byte(message), &data)
		if err != nil {
			client.Destroy()
			return
		}

		var userMap map[string]json.RawMessage

		err = json.Unmarshal([]byte(data["d"]), &userMap)
		if err != nil {
			client.Destroy()
			return
		}

		var userDatas []*LanyardData

		for _, item := range userMap {
			var userData *LanyardData

			err := json.Unmarshal(item, &userData)
			if err != nil {
				client.Destroy()
				return
			}

			userDatas = append(userDatas, userData)
		}

		presenceUpdate(userDatas)
	}

	client.socket.Connect()

	return client
}
