package lanyard

import (
	"encoding/json"
	"strings"
)

const (
	WS_URL = "wss://api.lanyard.rest/socket"
)

func singlePresenceUpdate(client WSClient, message string) (*LanyardData, error) {
	var data LanyardWSResponse
	err := json.Unmarshal([]byte(message), &data)
	if err != nil {
		return &LanyardData{}, err
	}

	return data.D, nil
}

func multiplePresenceUpdate(client WSClient, message string) ([]*LanyardData, error) {
	var data map[string]json.RawMessage
	err := json.Unmarshal([]byte(message), &data)
	if err != nil {
		return []*LanyardData{}, err
	}

	var userMap map[string]json.RawMessage
	err = json.Unmarshal([]byte(data["d"]), &userMap)
	if err != nil {
		return []*LanyardData{}, err
	}

	var userDatas []*LanyardData
	for _, item := range userMap {
		var userData *LanyardData
		err := json.Unmarshal(item, &userData)
		if err != nil {
			return []*LanyardData{}, err
		}

		userDatas = append(userDatas, userData)
	}

	return userDatas, nil
}

func ListenUser(userId string, presenceUpdate func(data *LanyardData)) WSClient {
	client := clientFactory("subscribe_to_id", "\"" + userId + "\"", func(client WSClient, message string) {
		data, err := singlePresenceUpdate(client, message)
		if err != nil {
			client.Destroy()
			return
		}

		presenceUpdate(data)
	})

	return *client
}

func ListenMultipleUsers(userIds []string, presenceUpdate func(data *LanyardData)) WSClient {
	var formattedIds []string
	for _, id := range userIds {
		formattedIds = append(formattedIds, "\""+id+"\"")
	}

	client := clientFactory("subscribe_to_ids", "[" + strings.Join(formattedIds, ",") + "]", func(client WSClient, message string) {
		if (strings.Contains(message, "INIT_STATE")) {
			userDatas, err := multiplePresenceUpdate(client, message)
			if err != nil {
				client.Destroy()
				return
			}

			for _, data := range userDatas {
				presenceUpdate(data)
			}
		} else {
			data, err := singlePresenceUpdate(client, message)
			if err != nil {
				client.Destroy()
				return
			}

			presenceUpdate(data)
		}
	})

	return *client
}
