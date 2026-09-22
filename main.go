package main

import (
	"fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Pokedex >")
        if scanner.Scan() == true {
            input_slice := cleanInput(scanner.Text())
            fmt.Printf("Your command was: %v", input_slice[0])
            fmt.Print("\n")
        }
    }
}


