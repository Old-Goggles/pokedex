package pokeapi

import (
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, cache *pokecache.Cache) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}
