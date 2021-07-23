package lanyard

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	API_URL = "https://api.lanyard.rest/v1/users/"
)

func FetchUser(userId string) (LanyardResponse, error) {
	resp, err := http.Get(API_URL + userId)

	// Check An Error for Send HTTP Request
	if err != nil {
		return LanyardResponse{
			false,
			&LanyardData{},
			&LanyardError{},
		}, err
	}

	body, err := io.ReadAll(resp.Body)

	// Check An Error for Read Body
	if err != nil {
		return LanyardResponse{
			false,
			&LanyardData{},
			&LanyardError{},
		}, err
	}

	var data LanyardResponse

	err = json.Unmarshal(body, &data)

	// Check Lanyard sent An Error
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return LanyardResponse{
			false,
			&LanyardData{},
			&LanyardError{},
		}, errors.New(data.Error.Message)
	}

	// Check An Error for JSON De-Convertion.
	if err != nil {
		return LanyardResponse{
			false,
			&LanyardData{},
			&LanyardError{},
		}, err
	}

	// Return Data with No Error
	return data, nil
}
