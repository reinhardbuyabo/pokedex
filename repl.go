package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	Next     *string
	Previous *string
}

type cliCommands struct {
	name        string
	description string
	callback    func(*Config) error
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			cfg := &Config{}
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMap(cfg *Config) error {
	url := ""

	if cfg.Next != nil {
		url = *cfg.Next
	}

	data, err := getNamesofLocations(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, locationArea := range data.Results {
		fmt.Println(locationArea.Name)
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err := getNamesofLocations(*cfg.Previous)
	if err != nil {
		return err
	}

	for _, locationArea := range data.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func getCommands() map[string]cliCommands {

	return map[string]cliCommands{
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
			description: "Displays the names of 20 locations in the Pokemon World",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of previoius 20 locations in the Pokemon World",
			callback:    commandMapb,
		},
	}
}
