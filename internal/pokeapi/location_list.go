package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreas, error) {

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if entry, exists := c.cache.Get(url); exists {
		fmt.Println("USED CACHE")
		pokeAreas := LocationAreas{}
		err := json.Unmarshal(entry, &pokeAreas)
		if err != nil {
			log.Fatal(err)
			return LocationAreas{}, err
		}
		return pokeAreas, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return LocationAreas{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return LocationAreas{}, err
	}

	pokeAreas := LocationAreas{}
	err = json.Unmarshal(body, &pokeAreas)
	if err != nil {
		log.Fatal(err)
		return LocationAreas{}, err
	}

	c.cache.Add(url, body)
	fmt.Println("DIDNT USED CACHE")

	return pokeAreas, nil
}
