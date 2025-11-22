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

type LocationFull struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func ExploreLocation(location string, pkc pokecache.Cache) (LocationFull, error) {

	locationPokemon := LocationFull{}

	url := "https://pokeapi.co/api/v2/location-area/" + location

	c, ok := pkc.Get(url)
	if ok {
		if err := json.Unmarshal(c.Val, &locationPokemon); err == nil {
			return locationPokemon, nil
		}
	}

	c, ok = pkc.Get(url)
	if ok {
		if err := json.Unmarshal(c.Val, &location); err == nil {
			return locationPokemon, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return locationPokemon, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationPokemon, err
	}

	if err := json.Unmarshal(data, &locationPokemon); err != nil {
		return locationPokemon, err
	}

	return locationPokemon, nil
}
