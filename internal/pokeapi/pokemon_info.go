package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) PokeInfo(name string) (Pokemon, error) {
	url := baseURL + fmt.Sprintf("/pokemon/%v", name)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Invalid pokemon")
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		log.Fatal(err)
		return Pokemon{}, err
	}

	return pokemon, nil

}
