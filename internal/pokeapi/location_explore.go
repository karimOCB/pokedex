package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationExplore(parameters string) (LocationEncounters, error) {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	actualUrl := baseUrl + parameters
	encounters := LocationEncounters{}

	val, ok := c.cache.Get(actualUrl)

	if ok {
		if err := json.Unmarshal(val, &encounters); err != nil {
			return encounters, err
		}
		return encounters, nil
	}

	req, err := http.NewRequest("GET", actualUrl, nil)

	if err != nil {
		return encounters, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return encounters, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return encounters, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return encounters, err
	}

	if err := json.Unmarshal(data, &encounters); err != nil {
		return encounters, err
	}

	c.cache.Add(actualUrl, data)

	return encounters, nil
}
