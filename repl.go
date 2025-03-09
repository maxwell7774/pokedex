package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
        if len(words) == 0 {
            continue
        }

		command, exists := getCommands()[words[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
        err := command.callback()
        if err != nil {
            fmt.Println(err)
        }
	}
}


func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}


type cliCommand struct {
	name        string
	description string
	callback    func() error
}


func getCommands() map[string]cliCommand {
    return map[string]cliCommand {
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
    }
}
