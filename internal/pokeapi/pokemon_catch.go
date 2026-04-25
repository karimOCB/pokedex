package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonCatch(pokemonName string) (Pokemon, error) {
	baseURL := "https://pokeapi.co/api/v2/pokemon/"
	actualURL := baseURL + pokemonName
	var pokemonInfo Pokemon

	val, ok := c.cache.Get(actualURL)
	if ok {
		if err := json.Unmarshal(val, &pokemonInfo); err != nil {
			return pokemonInfo, err
		}
		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", actualURL, nil)
	if err != nil {
		return pokemonInfo, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return pokemonInfo, nil
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return pokemonInfo, fmt.Errorf("Failed res with status code: %d", res.StatusCode)
	}

	val, err = io.ReadAll(res.Body)
	if err != nil {
		return pokemonInfo, err
	}

	if err := json.Unmarshal(val, &pokemonInfo); err != nil {
		return pokemonInfo, err
	}

	c.cache.Add(actualURL, val)

	return pokemonInfo, nil
}
