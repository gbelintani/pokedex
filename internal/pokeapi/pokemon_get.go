package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	finalUrl := baseURL + "/pokemon/" + name

	cached, exists := c.cache.Get(finalUrl)
	if exists {
		var apiResponse Pokemon
		if err := json.Unmarshal(cached, &apiResponse); err != nil {
			return Pokemon{}, fmt.Errorf("could not unmarshal %w", err)
		}
		fmt.Println("from cache")
		return apiResponse, nil
	}
	res, err := c.client.Get(finalUrl)
	if err != nil {
		return Pokemon{}, fmt.Errorf("could not get: %w", err)
	}
	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("could not get: %s", res.Status)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading data: %w", err)
	}

	var apiResponse Pokemon
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return Pokemon{}, fmt.Errorf("error decoding: %w", err)
	}

	c.cache.Add(finalUrl, data)

	return apiResponse, nil
}
