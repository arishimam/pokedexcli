package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListPokemon(name string) (LocationDetails, error) {

	url := baseURL + fmt.Sprintf("/location-area/%v", name)

	if entry, exists := c.cache.Get(url); exists {
		fmt.Println("USED CACHE")
		locationDetails := LocationDetails{}
		err := json.Unmarshal(entry, &locationDetails)
		if err != nil {
			log.Fatal(err)
			return LocationDetails{}, err
		}
		return locationDetails, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return LocationDetails{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return LocationDetails{}, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return LocationDetails{}, nil
	}

	locationDetails := LocationDetails{}
	err = json.Unmarshal(body, &locationDetails)
	if err != nil {
		log.Fatal(err)
		return LocationDetails{}, nil
	}

	c.cache.Add(url, body)
	fmt.Println("DIDNT USED CACHE")

	return locationDetails, nil
}
