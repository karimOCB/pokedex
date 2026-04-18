package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	Client http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		Client: http.Client{Timeout: timeout},
	}
}