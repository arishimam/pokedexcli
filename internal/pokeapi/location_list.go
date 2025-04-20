package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/arishimam/pokedexcli/internal/cache"
)

func (c *Client) ListLocations(C *cache.Cache, pageURL *string) (LocationAreas, error) {

	pokeAreas := LocationAreas{}

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	entry, exists := c.cache.Get(url)
	if exists {

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

	// unmarshall data into struct
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return LocationAreas{}, err
	}

	err = json.Unmarshal(body, &pokeAreas)
	if err != nil {
		log.Fatal(err)
		return LocationAreas{}, err
	}

	c.cache.Add(url, body)

	return pokeAreas, nil
}
