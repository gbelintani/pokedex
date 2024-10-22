package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gbelintani/pokedex/internal/pokecache"
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

func (c *Client) GetLocationAreas(url *string) (PokeApiListLocationResponse, error) {
	finalUrl := "https://pokeapi.co/api/v2/location-area"
	if url != nil && len(*url) > 0 {
		finalUrl = *url
	}

	cached, exists := c.cache.Get(finalUrl)
	if exists {
		var apiResponse PokeApiListLocationResponse
		if err := json.Unmarshal(cached, &apiResponse); err != nil {
			return PokeApiListLocationResponse{}, fmt.Errorf("could not unmarshal %w", err)
		}
		fmt.Println("from cache")
		return apiResponse, nil
	}

	res, err := c.client.Get(finalUrl)
	if err != nil {
		return PokeApiListLocationResponse{}, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeApiListLocationResponse{}, fmt.Errorf("error reading data: %w", err)
	}

	var apiResponse PokeApiListLocationResponse
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return PokeApiListLocationResponse{}, fmt.Errorf("error decoding: %w", err)
	}

	c.cache.Add(finalUrl, data)

	return apiResponse, nil
}
