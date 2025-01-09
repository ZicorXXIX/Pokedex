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

func NewClient(timeout time.Duration) Client {
    c := pokecache.NewCache(100 * time.Second)
    return Client{
        httpClient: http.Client{
            Timeout: timeout,
        },
        cache: c,
    }
}
