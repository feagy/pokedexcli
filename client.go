package main

import (
    "net/http"
    "encoding/json"
    "io"
    "fmt"
)

type location struct {
    Name string `json:"name"`
    Url string `json:"url"`
}

type locations struct {
    Count int `json:"count"`
    Next string `json:"next"`
    Previous string `json:"previous"`
    Results []location `json:"results"`
}

type encounters struct {
    PokemonEncounters [] struct {
        Pokemon struct {
            Name string `json:"name"`
            Url string `json:"url"`
        } `json:"pokemon"`
    } `json:"pokemon_encounters"`
}

type pokemon struct {
    Name string `json:"name"`
    Height int `json:"height"`
    Weight int `json:"weight"`
    BaseExperience int `json:"base_experience"`
    Stats []struct {
        BaseStat int `json:"base_stat"`
        Stat struct {
            Name string `json:"name"`
            Url string `json:"url"`
        } `json:"stat"`
    } `json:"stats"`
    Types []struct {
        Type struct{
            Name string `json:"name"`
            Url string `json:"url"`
        } `json:"type"`
    } `json:"types"`
}

func getLocations(url string, config *config) (locations, error)  {
    dat, ok := config.cache.Get(url)
    if !ok {
        res, err := http.Get(url)
        if err != nil {
            return locations{}, err
        }
        defer res.Body.Close()
            
	    new_dat, err := io.ReadAll(res.Body)
	    if err != nil {
		    return locations{}, err
	    }
        dat = new_dat
        config.cache.Add(url, dat)
    }

    lcts := locations{}
	err := json.Unmarshal(dat, &lcts)
	if err != nil {
		return locations{}, err
	}

    return lcts, nil
}

func getEncounters(url string, config *config) (encounters, error) {
    dat, ok := config.cache.Get(url)
    if !ok {
        res, err := http.Get(url)
        if err != nil {
            return encounters{}, err
        }
        defer res.Body.Close()
        if res.StatusCode == 404 {
            return encounters{}, fmt.Errorf("Location not found")
        }

	    new_dat, err := io.ReadAll(res.Body)
	    if err != nil {
		    return encounters{}, err
	    }
        dat = new_dat
        config.cache.Add(url, dat)
    }
    enctrs := encounters{}
	err := json.Unmarshal(dat, &enctrs)
	if err != nil {
		return encounters{}, err
	}
    return enctrs, nil
}

func getPokemon(url string, config *config) (pokemon, error) {
    dat, ok := config.cache.Get(url)
    if !ok {
        res, err := http.Get(url)
        if err != nil {
            return pokemon{}, err
        }
        defer res.Body.Close()
        if res.StatusCode == 404 {
            return pokemon{}, fmt.Errorf("Pokemon not found")
        }

	    new_dat, err := io.ReadAll(res.Body)
	    if err != nil {
		    return pokemon{}, err
	    }
        dat = new_dat
        config.cache.Add(url, dat)
    }
    pkmn := pokemon{}
	err := json.Unmarshal(dat, &pkmn)
	if err != nil {
		return pokemon{}, err
	}
    return pkmn, nil
}
