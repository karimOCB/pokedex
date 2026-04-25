package pokeapi

type Locations struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationEncounters struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
