package pokeapi

import (
    "net/http"
	"encoding/json"
    "io"
)


func (c *Client) GetLocations(pageUrl *string) (LocationAreaResponse, error) {
    url := baseUrl +  "/location-area"
    if pageUrl != nil {
        url = *pageUrl
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return LocationAreaResponse{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return LocationAreaResponse{}, err
    }

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return LocationAreaResponse{}, err
    }
    defer res.Body.Close()

    var data LocationAreaResponse
    if json.Valid([]byte(body)){
        json.Unmarshal([]byte(body), &data)
    }

    return data, nil
}
