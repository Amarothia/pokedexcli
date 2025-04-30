package main

import (
	"fmt"

	"github.com/Amarothia/pokedexcli/funcs"
)

func main() {
	fmt.Println("Hello, World!")
	result := funcs.CleanInput("Hello World!")
	fmt.Println(result)
}
