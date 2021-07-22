package lanyard

import (
	"encoding/json"
	"io"
	"net/http"
)


const (
	API_URL = "https://api.lanyard.rest/v1/users/"
)

func FetchUser(userId string) LanyardResponse {
	resp, err := http.Get(API_URL + userId)
	
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data LanyardResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	return data
}
