package funcs

import (
	"fmt"
	"math/rand"
	"os"
)

const pokeapiURL = "https://pokeapi.co/api/v2/"

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

type Pokemon struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	XP     int    `json:"base_experience"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Types  []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
}

type Config struct {
	Next              *string
	Previous          *string
	LocationAreaCache map[string]LocationAreas
	PokemonCache      map[string]LocationAreaEncounters
	Player            Player
	AreaName          *string
	PokemonName       *string
	InspectPokemon    *string
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
		"catch": {
			name:        "catch",
			description: "Catch <pokemon name> to try to catch a Pokemon",
			callback:    CommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect <pokemon name> to learn more about a Pokemon",
			callback:    CommandInspect,
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

func CommandInspect(cfg *Config) error {

	pokemon, ok := cfg.Player.InspectPokemon(*cfg.InspectPokemon)

	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n  -hp: %d\n  -attack: %d\n  -defense %d\n  -special-attack %d\n  -special-defense %d\n  -speed %d\nTypes:\n",
		pokemon.Name, pokemon.Height, pokemon.Weight, pokemon.Stats[0].BaseStat, pokemon.Stats[1].BaseStat, pokemon.Stats[2].BaseStat, pokemon.Stats[3].BaseStat, pokemon.Stats[4].BaseStat, pokemon.Stats[5].BaseStat)

	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}

	return nil
}

func CommandCatch(cfg *Config) error {
	url := pokeapiURL + "pokemon/" + *cfg.PokemonName

	pokemon, err := GetPokemon(url)
	if err != nil {
		fmt.Printf("You have failed to name that Pokemon!\n")
		return nil
	}

	chance := rand.Intn(500)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if pokemon.XP <= chance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.Player.AddPokemon(pokemon)
	} else {
		fmt.Printf("%s escaped! (Xp: %d > Chance: %d)\n", pokemon.Name, pokemon.XP, chance)
	}

	return nil

}

func CommandExplore(cfg *Config) error {
	if cfg.AreaName == nil {
		return fmt.Errorf("explore syntax incorrect, please use explore <area name>")
	}

	url := pokeapiURL + "location-area/" + *cfg.AreaName

	var lae LocationAreaEncounters
	var ok bool
	var err error

	if lae, ok = cfg.PokemonCache[url]; !ok {
		lae, err = GetLocationAreaEncounters(url)
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
	url := pokeapiURL + "location-area/" + "?offset=0&limit=20"

	if cfg.Next != nil {
		url = *cfg.Next
	}

	var areaList LocationAreas
	var ok bool
	var err error

	if areaList, ok = cfg.LocationAreaCache[url]; !ok {
		areaList, err = GetLocationAreas(url)
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
	url := pokeapiURL + "location-area/" + "?offset=0&limit=20"

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
