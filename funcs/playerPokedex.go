package funcs

import (
	"fmt"
	"sort"
)

type Player struct {
	CaughtPokemon map[string]Pokemon
	Pokeballs     int
}

func (P Player) AddPokemon(p Pokemon) {
	P.CaughtPokemon[p.Name] = p
}

func (P Player) InspectPokemon(name string) (Pokemon, bool) {
	if pokemon, ok := P.CaughtPokemon[name]; ok {
		return pokemon, true
	}
	return Pokemon{}, false
}

func (P Player) GetPokedex() error {
	if len(P.CaughtPokemon) == 0 {
		return fmt.Errorf("you have no pokemon yet")
	}

	fmt.Println("Your Pokedex:")

	keys := make([]string, 0, len(P.CaughtPokemon))
	for key := range P.CaughtPokemon {
		keys = append(keys, key)
	}

	// Sort the keys
	sort.Strings(keys)

	// Iterate through the map in sorted key order
	for _, key := range keys {
		pokemon := P.CaughtPokemon[key]
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
