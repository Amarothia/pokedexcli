package funcs

import (
	"bufio"
	"fmt"
	"os"
)

func StartRepl() {

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := CleanInput(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		fmt.Printf("Your command was: %s\n", commandName)
	}

}
