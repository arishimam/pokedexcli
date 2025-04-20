package pokeapi

import (
	"github.com/arishimam/pokedexcli/internal/cache"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      cache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache.NewCache(cacheInterval),
	}
}
