package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/christophergantes/pokedexcli/internal/pokeapi"
)

type CLICommand struct {
	name        string
	description string
	callback    func(*Config) error
	config      struct {
		Next     string
		Previous string
	}
}

type Config struct {
	Next     *string
	Previous *string
	Client   *pokeapi.Client
}

func GetCommands() map[string]CLICommand {
	return map[string]CLICommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit this Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays all pokemon location areas and proceeds to next page if called multiple times",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "returns to previous page of map if called",
			callback:    commandMapB,
		},
	}
}

func cleanInput(text string) []string {
	s := strings.ToLower(text)
	s = strings.TrimSpace(s)
	return strings.Fields(s)
}

func commandHelp(config *Config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(config *Config) error {
	var url *string = nil
	if config.Next != nil {
		url = config.Next
	}
	resources, err := config.Client.GetLocationAreas(url)
	if err != nil {
		return err
	}
	config.Next = resources.Next
	config.Previous = resources.Previous

	for _, result := range resources.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapB(config *Config) error {
	var url *string = nil
	if config.Previous != nil {
		url = config.Previous
	}
	resources, err := config.Client.GetLocationAreas(url)
	if err != nil {
		return err
	}
	config.Next = resources.Next
	config.Previous = resources.Previous

	for _, result := range resources.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func NewConfig() *Config {
	return &Config{
		Client:   pokeapi.NewClient(),
		Next:     nil,
		Previous: nil,
	}
}
