package pokeapi

import (
	"net/http"
	"time"

	"github.com/ZicorXXIX/pokedex/internal/pokecache"
)

type Client struct {
    httpClient http.Client
    cache      *pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
    c := pokecache.NewCache(cacheInterval)
    return Client{
        httpClient: http.Client{
            Timeout: timeout,
        },
        cache: c,
    }
}
