package pokeapi

import (
	"net/http"
	"time"

	"github.com/karimOCB/pokedex/internal/pokecache"
)

type Client struct {
	Client http.Client
	cache  *pokecache.Cache
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		Client: http.Client{Timeout: timeout},
		cache:  pokecache.NewCache(5 * time.Second),
	}
}
