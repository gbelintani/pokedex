package pokeapi

import (
	"net/http"
	"time"

	"github.com/gbelintani/pokedex/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Client struct {
	client http.Client
	cache  pokecache.Cache
}

func NewClient() Client {
	client := Client{
		client: http.Client{},
		cache:  pokecache.NewCache(10 * time.Second),
	}
	return client
}
