package main

import (
    "net/http"
    "encoding/json"
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


func getLocations(url string) (locations, error)  {
    res, err := http.Get(url)
    if err != nil {
        return locations{}, err
    }
    defer res.Body.Close()

    lcts := locations{}
    decoder := json.NewDecoder(res.Body)
    if err := decoder.Decode(&lcts); err != nil {
	    return locations{}, err
	}

    return lcts, nil
}
