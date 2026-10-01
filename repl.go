package main

import (
    "strings"
    "bufio"
    "os"
    "fmt"
    "math/rand"
)

func cleanInput(text string) []string {
    return strings.Fields(strings.TrimSpace(strings.ToLower(text)))
}

func startREPL(config *config){
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("\nPokedex >")
        if scanner.Scan() == true {
            input_slice := cleanInput(scanner.Text())
            command, ok := config.commands[input_slice[0]]
            if !ok {
                fmt.Printf("Unknown command: %v\n", input_slice[0])
                continue
            }
            err := command.callback(config, input_slice[1:])
            if err != nil {
                fmt.Printf("An error occured: %v\n", err)
            }
        }
    }
}

func commandExit(config *config, arguments []string) error {
    fmt.Printf("Closing the Pokedex... Goodbye!\n")
    os.Exit(0)
    return nil
}

func commandHelp(config *config, arguments []string) error {
    fmt.Println("Welcome to the Pokedex!")
    for _, v := range config.commands {
        fmt.Printf("Command: %v\n", v.name)
        fmt.Printf("Description: %v\n\n", v.description)
    }
    return nil
}

func commandMap(config *config, arguments []string) error {
    url := config.next

    locations, err := getLocations(url, config)

    if err != nil {
        return err    
    }
    
    config.next = locations.Next
    config.previous = locations.Previous    

    for _, location := range locations.Results {
        fmt.Println("- " + location.Name)
    }

    return nil
}

func commandMapb(config *config, arguments []string) error {
    url := config.previous
    if url == "" {
        fmt.Println("you're on the first page")
        return nil
    }
    locations, err := getLocations(url, config)
    if err != nil {
        return err    
    }
    
    config.next = locations.Next
    config.previous = locations.Previous    

    for _, location := range locations.Results {
        fmt.Println("- " + location.Name)
    }

    return nil
}

func commandExplore(config *config, arguments []string) error {
    if len(arguments) == 0 {
        fmt.Println("you should provide the name of the area you want to explore")
        return nil
    }
    url := config.explore_url + arguments[0]
    fmt.Printf("Exploring %v\n", arguments[0])
    encounters, err := getEncounters(url, config)
    if err != nil {
        return err    
    }
    
    fmt.Println("Found Pokemon:")
    for _, encounter := range encounters.PokemonEncounters {
        fmt.Println("- " + encounter.Pokemon.Name)
    }

    return nil
}

func commandCatch(config *config, arguments []string) error {
    if len(arguments) == 0 {
        fmt.Println("you should provide the name of the pokemon you want to catch")
        return nil
    }

    url := config.pokemon_url + arguments[0]
    pokemon, err := getPokemon(url, config)
    if err != nil {
        return err    
    }
    
    fmt.Printf("Throwing a Pokeball at %v...\n", arguments[0])
    res := rand.Intn(pokemon.BaseExperience)
    if res <= config.catching_threshold {
        fmt.Printf("%v was caught!\n", arguments[0])
        config.pokedex[arguments[0]] = pokemon
    } else {
        fmt.Printf("%v escaped!\n", arguments[0])
    }
    return nil
}

func commandInspect(config *config, arguments []string) error {
    if len(arguments) == 0 {
        fmt.Println("you should provide the name of the pokemon you want to inspect")
        return nil
    }

    pokemon, ok := config.pokedex[arguments[0]]
    if !ok {
        fmt.Println("you have not caught that pokemon")
        return nil
    }
    fmt.Printf("Name: %v\nHeight:%v\nWeight:%v\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
    for _, s := range pokemon.Stats {
        fmt.Printf(" -%v: %v\n", s.Stat.Name, s.BaseStat)
    }
    fmt.Println("Types:")
    for _, t := range pokemon.Types {
        fmt.Printf(" - %v\n", t.Type.Name)
    }
    return nil
}

func commandPokedex(config *config, arguments []string) error {
    if len(config.pokedex) == 0 {
        fmt.Println("you have not caught any pokemon yet")
        return nil
    }
    fmt.Println("Your Pokedex:")
    for name, _ := range config.pokedex {
        fmt.Printf(" - %v\n", name)
    }
    return nil
}
