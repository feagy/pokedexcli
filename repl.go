package main

import (
    "strings"
    "bufio"
    "os"
    "fmt"
)

func cleanInput(text string) []string {
    return strings.Fields(strings.TrimSpace(strings.ToLower(text)))
}

func startREPL(config *config){
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex >")
        if scanner.Scan() == true {
            input_slice := cleanInput(scanner.Text())
            command, ok := config.commands[input_slice[0]]
            if !ok {
                fmt.Printf("Unknown command: %v", input_slice[0])
            }
            err := command.callback(config)
            if err != nil {
                fmt.Errorf("An erroe occured: %v", err)
            }
        }
    }
}

func commandExit(config *config) error {
    fmt.Printf("Closing the Pokedex... Goodbye!\n")
    os.Exit(0)
    return nil
}

func commandHelp(config *config) error {
    fmt.Println("Welcome to the Pokedex!\n")
    for _, v := range config.commands {
        fmt.Printf("Command: %v\n", v.name)
        fmt.Printf("Description: %v\n\n", v.description)
    }
    return nil
}
