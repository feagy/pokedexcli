package main

import (
    "strings"
    "bufio"
    "os"
    "fmt"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands( )map[string]cliCommand{
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
    }
}

func cleanInput(text string) []string {
    return strings.Fields(strings.TrimSpace(strings.ToLower(text)))
}

func startREPL(){
    scanner := bufio.NewScanner(os.Stdin)
    
    for {
        fmt.Print("Pokedex >")
        if scanner.Scan() == true {
            input_slice := cleanInput(scanner.Text())
            command, ok := getCommands()[input_slice[0]]
            if !ok {
                fmt.Printf("Unknown command: %v", input_slice[0])
            }
            err := command.callback()
            if err != nil {
                fmt.Errorf("An erroe occured: %v", err)
            }
        }
    }
}

func commandExit() error {
    fmt.Printf("Closing the Pokedex... Goodbye!\n")
    os.Exit(0)
    return nil
}

func commandHelp() error {
    fmt.Println("Welcome to the Pokedex!\n")
    for _, v := range getCommands() {
        fmt.Printf("Command: %v\n", v.name)
        fmt.Printf("Description: %v\n\n", v.description)
    }
    return nil
}
