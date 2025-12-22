package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	var pokemon Pokemon

	if data, ok := c.cache.Get(url); ok {
		json.Unmarshal(data, &pokemon)
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
