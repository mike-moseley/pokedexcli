package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/location-area"

type Client struct {
	httpClient http.Client
}

func NewClient(duration time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: duration,
		},
	}
}

func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	if pageURL == nil {
		newURL := baseURL
		pageURL = &newURL
	}
	res, err := c.httpClient.Get(*pageURL)
	if err != nil {
		return RespLocations{}, fmt.Errorf("Error fetching from pokeapi: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocations{}, fmt.Errorf("Error reading response from pokeapi: %w", err)
	}

	var resp_locations RespLocations
	err = json.Unmarshal(body, &resp_locations)
	if err != nil {
		return RespLocations{}, fmt.Errorf("Error unmarshaling json: %w", err)
	}

	return resp_locations, nil
}
