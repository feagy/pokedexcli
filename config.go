package main

import (
    "time"
)

type config struct{
    commands map[string]cliCommand
    next     string
    previous string
    explore_url string
    pokemon_url string
    pokedex map[string]pokemon
    catching_threshold int
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
        "catch": {
            name:        "catch <pokemon_name>",
            description: "Try to catch the pokemon",
            callback:    commandCatch,
        },
        "inspect": {
            name:        "inspect <pokemon_name>",
            description: "See the details about a caught pokemon",
            callback:    commandInspect,
        },
        "pokedex": {
            name:        "pokedex",
            description: "List all the pokemon you caught",
            callback:    commandPokedex,
        },
    }
}

func getConfig() *config {
    return &config{
        commands: getCommands(),
        next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
        previous: "",
        explore_url: "https://pokeapi.co/api/v2/location-area/",
        pokemon_url: "https://pokeapi.co/api/v2/pokemon/",
        pokedex: map[string]pokemon{},
        catching_threshold: 40,
        cache: newCache(10 * time.Second),  
    }
}
