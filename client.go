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


func getLocations(url string, config *config) (locations, error)  {
    dat, ok := config.cache.Get(url)
    if !ok {
        fmt.Println("cache not used")
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
    } else {
        fmt.Println("cache used")
    }

    lcts := locations{}
	err := json.Unmarshal(dat, &lcts)
	if err != nil {
		return locations{}, err
	}

    return lcts, nil
}
