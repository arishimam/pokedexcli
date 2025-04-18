package pokeapi

import (
	"encoding/json"
	"fmt"
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

	fmt.Println("Getting: ", url)
	entry, exists := C.Get(url)
	if exists {
		fmt.Println("USED THE CACHE")
		fmt.Println()

		err := json.Unmarshal(entry, &pokeAreas)
		if err != nil {
			log.Fatal(err)
			return LocationAreas{}, err
		}
		return pokeAreas, nil
	}

	fmt.Println("DIDNT USE THE CACHE")
	fmt.Println()
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

	C.Add(url, body)
	fmt.Println("Adding: ", url)
	err = json.Unmarshal(body, &pokeAreas)
	if err != nil {
		log.Fatal(err)
		return LocationAreas{}, err
	}

	return pokeAreas, nil
}
