package funcs

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
