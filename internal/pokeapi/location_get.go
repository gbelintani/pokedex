package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocationDetail(location string) (PokeApiDeepLocationResponse, error) {
	finalUrl := baseURL + "/location-area/" + location
	cached, exists := c.cache.Get(finalUrl)
	if exists {
		var apiResponse PokeApiDeepLocationResponse
		if err := json.Unmarshal(cached, &apiResponse); err != nil {
			return PokeApiDeepLocationResponse{}, fmt.Errorf("could not unmarshal %w", err)
		}
		fmt.Println("from cache")
		return apiResponse, nil
	}

	res, err := c.client.Get(finalUrl)
	if err != nil {
		return PokeApiDeepLocationResponse{}, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeApiDeepLocationResponse{}, fmt.Errorf("error reading data: %w", err)
	}

	var apiResponse PokeApiDeepLocationResponse
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		return PokeApiDeepLocationResponse{}, fmt.Errorf("error decoding: %w", err)
	}

	c.cache.Add(finalUrl, data)

	return apiResponse, nil

}
