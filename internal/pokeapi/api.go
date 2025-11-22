package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vimto1234/pokedexcli/internal/pokecache"
)

type LocationResult struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	Url  string `json:url`
}

func GetLocation(url string, pkc pokecache.Cache) (LocationResult, error) {

	location := LocationResult{}

	c, ok := pkc.Get(url)
	if ok {
		if err := json.Unmarshal(c.Val, &location); err == nil {
			return location, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return location, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return location, err
	}

	if err := json.Unmarshal(data, &location); err != nil {
		return location, err
	}

	return location, nil
}
