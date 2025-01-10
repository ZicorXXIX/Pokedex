package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(pokemonName string) (PokemonDetails, error) {
	url := baseUrl + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := PokemonDetails{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return PokemonDetails{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetails{}, err
	}

	pokemonResp := PokemonDetails{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonDetails{}, err
	}

	c.cache.Add(url, dat)

	return pokemonResp, nil
}
