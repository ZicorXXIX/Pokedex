package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(name string) (PokemonDetails, error) {
    url := baseUrl + "/pokemon/" + name

    var data PokemonDetails
    cachedData, ok := c.cache.Get(url)
    if ok {
        if err:=json.Unmarshal(cachedData, &data); err != nil {
            return PokemonDetails{}, err
        }
        return data, nil
    }

    req, err := http.NewRequest("GET", url, nil) 
    if err != nil {
        return PokemonDetails{}, nil
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return PokemonDetails{}, err
    }

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return PokemonDetails{}, err
    }
    defer res.Body.Close()

    if json.Valid([]byte(body)) {
        json.Unmarshal([]byte(body), &data)
    }

    c.cache.Add(url, []byte(body))

    return data, nil    
}
