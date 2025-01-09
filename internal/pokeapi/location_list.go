package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (c *Client) GetLocations(pageUrl *string) (LocationAreaResponse, error) {
    url := baseUrl +  "/location-area"
    if pageUrl != nil {
        url = *pageUrl
    }
    cachedData, exists := c.cache.Get(url)

    var data LocationAreaResponse
    if exists {
        fmt.Println("Retreiving Cached Data")
        if err:= json.Unmarshal(cachedData, &data); err != nil {
            return LocationAreaResponse{}, err
        }
        return data, nil
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
    c.cache.Add(url, []byte(body))

    if json.Valid([]byte(body)){
        json.Unmarshal([]byte(body), &data)
    }


    return data, nil
}
