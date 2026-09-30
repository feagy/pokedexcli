package main

import (
    "time"
)

type config struct{
    commands map[string]cliCommand
    next     string
    previous string
    explore_url string
    cache *cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands()map[string]cliCommand{
    return map[string]cliCommand{
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "map": {
            name:        "map",
            description: "Displays the name of locations",
            callback:    commandMap,
        },
        "mapb": {
            name:        "mapb",
            description: "Displays the name of locations at the previous page",
            callback:    commandMapb,
        }, 
        "explore": {
            name:        "explore <area_name>",
            description: "Explore the are for pokemons",
            callback:    commandExplore,
        },
    }
}

func getConfig() *config {
    return &config{
        commands: getCommands(),
        next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
        previous: "",
        explore_url: "https://pokeapi.co/api/v2/location-area/",
        cache: newCache(10 * time.Second),  
    }
}
