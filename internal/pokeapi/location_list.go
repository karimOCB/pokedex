package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

func (c *Client) ListLocations(url *string) (Locations, error) {
	var actualURL string
	var locations Locations

	if url == nil {
		actualURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		actualURL = *url
	}
	
	req, err := http.NewRequest("GET", actualURL, nil)
	if err != nil {
		return locations, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return locations, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return locations, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locations, err
	}
	
	if err := json.Unmarshal(data, &locations); err != nil {
		return locations, err
	}

	return locations, nil
}