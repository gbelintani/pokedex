package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLocationAreas(url *string) (PokeApiListLocationResponse, error) {
	finalUrl := "https://pokeapi.co/api/v2/location-area"
	if url != nil && len(*url) > 0 {
		finalUrl = *url
	}
	res, err := http.Get(finalUrl)
	if err != nil {
		return PokeApiListLocationResponse{}, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	var apiResponse PokeApiListLocationResponse
	if err := decoder.Decode(&apiResponse); err != nil {
		return PokeApiListLocationResponse{}, fmt.Errorf("error decoding: %w", err)
	}

	return apiResponse, nil
}
