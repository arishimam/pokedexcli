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

	return locationDetails, nil
}
