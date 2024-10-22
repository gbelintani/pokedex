package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokeApiResponse struct {
	Count    int                     `json:"count"`
	Next     *string                 `json:"next"`
	Previous *string                 `json:"previous"`
	Results  []PokeApiResultResponse `json:"results"`
}

type PokeApiResultResponse struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocationAreas(url *string) (PokeApiResponse, error) {
	finalUrl := "https://pokeapi.co/api/v2/location-area"
	if url != nil && len(*url) > 0 {
		finalUrl = *url
	}
	res, err := http.Get(finalUrl)
	if err != nil {
		return PokeApiResponse{}, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	var apiResponse PokeApiResponse
	if err := decoder.Decode(&apiResponse); err != nil {
		return PokeApiResponse{}, fmt.Errorf("error decoding: %w", err)
	}

	return apiResponse, nil
}
