package funcs

import (
	"fmt"
	"os"
)

const pokeapiURL = "https://pokeapi.co/api/v2/location-area/"

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

type LocationAreaEncounters struct {
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Config struct {
	Next              *string
	Previous          *string
	LocationAreaCache map[string]LocationAreas
	PokemonCache      map[string]LocationAreaEncounters
	AreaName          *string
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
		"explore": {
			name:        "explore",
			description: "Explore <area name> in order to see the Pokemon that live there",
			callback:    CommandExplore,
		},
	}

	return commands
}

func CommandExplore(cfg *Config) error {
	if cfg.AreaName == nil {
		return fmt.Errorf("explore syntax incorrect, please use explore <area name>")
	}

	url := pokeapiURL + *cfg.AreaName

	var lae LocationAreaEncounters
	var ok bool
	var err error

	if lae, ok = cfg.PokemonCache[url]; !ok {
		lae, err = GetPokemon(url)
		if err != nil {
			return fmt.Errorf("encountered error: %v", err)
		}
	}

	for _, encounter := range lae.Encounters {
		fmt.Printf("-%s\n", encounter.Pokemon.Name)
	}

	cfg.PokemonCache[url] = lae

	return nil
}

func CommandMap(cfg *Config) error {
	url := pokeapiURL + "?offset=0&limit=20"

	if cfg.Next != nil {
		url = *cfg.Next
	}

	var areaList LocationAreas
	var ok bool
	var err error

	if areaList, ok = cfg.LocationAreaCache[url]; !ok {
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
	cfg.LocationAreaCache[url] = areaList

	return nil
}

func CommandMapBack(cfg *Config) error {
	url := pokeapiURL + "?offset=0&limit=20"

	if cfg.Previous != nil {
		url = *cfg.Previous
	}

	areaList := cfg.LocationAreaCache[url]

	for _, area := range areaList.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = areaList.Next
	cfg.Previous = areaList.Previous

	return nil
}
