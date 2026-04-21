package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(url *string) (Locations, error) {
	var actualURL string
	var locations Locations

	if url == nil {
		actualURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		actualURL = *url
	}

	val, ok := c.cache.Get(actualURL)
	if ok {
		if err := json.Unmarshal(val, &locations); err != nil {
			return locations, err
		}
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

	c.cache.Add(actualURL, data)

	return locations, nil
}
