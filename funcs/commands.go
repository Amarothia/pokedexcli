package funcs

import (
	"fmt"
	"os"
)

const pokeapiURL = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"

type CliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Next     *string
	Previous *string
	Cache    map[string]LocationAreas
}

func CommandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(cfg *Config) error {
	commands := GetCommands()

	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func GetCommands() map[string]CliCommand {

	commands := map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "Explains the functions of the Pokedex",
			callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Shows the next 20 locations of the PokeWorld",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 locations of the PokeWorld",
			callback:    CommandMapBack,
		},
	}

	return commands
}

func CommandMap(cfg *Config) error {
	url := pokeapiURL

	if cfg.Next != nil {
		url = *cfg.Next
	}

	var areaList LocationAreas
	var ok bool
	var err error

	if areaList, ok = cfg.Cache[url]; !ok {
		areaList, err = GetLocations(url)
		if err != nil {
			return fmt.Errorf("error encountered: %v", err)
		}
	}

	for _, area := range areaList.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = areaList.Next
	cfg.Previous = areaList.Previous
	cfg.Cache[url] = areaList

	return nil
}

func CommandMapBack(cfg *Config) error {
	url := pokeapiURL

	if cfg.Previous != nil {
		url = *cfg.Previous
	}

	/*
		var areaList LocationAreas
		var ok bool
		var err error

		if areaList, ok = cfg.Cache[url]; !ok {
			areaList, err = GetLocations(url)
			if err != nil {
				return fmt.Errorf("error encountered: %v", err)
			}
		}
	*/

	areaList := cfg.Cache[url]

	for _, area := range areaList.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = areaList.Next
	cfg.Previous = areaList.Previous

	return nil
}
