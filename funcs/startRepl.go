package funcs

import (
	"bufio"
	"fmt"
	"os"
)

func StartRepl() {
	cfg := &Config{}
	cfg.LocationAreaCache = make(map[string]LocationAreas)
	cfg.PokemonCache = make(map[string]LocationAreaEncounters)

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := CleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if commandName == "explore" && len(words) < 2 {
			fmt.Println("Please use syntax 'explore <area name>' to find Pokemon!")
			continue
		} else if commandName == "explore" {
			cfg.AreaName = &words[1]
		}

		command, exists := GetCommands()[commandName]

		if exists {
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
