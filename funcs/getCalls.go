package funcs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLocationAreas(url string) (LocationAreas, error) {

	var areaList LocationAreas

	res, err := http.Get(url)
	if err != nil {
		return areaList, fmt.Errorf("encountered error: %v", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&areaList)
	if err != nil {
		return areaList, fmt.Errorf("encountered error: %v", err)
	}

	return areaList, nil
}

func GetLocationAreaEncounters(url string) (LocationAreaEncounters, error) {
	var lae LocationAreaEncounters

	res, err := http.Get(url)
	if err != nil {
		return lae, fmt.Errorf("encountered error: %v", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&lae)
	if err != nil {
		return lae, fmt.Errorf("encountered error: %v", err)
	}

	return lae, nil
}

func GetPokemon(url string) (Pokemon, error) {
	var pokemon Pokemon

	res, err := http.Get(url)
	if err != nil {
		return pokemon, fmt.Errorf("encountered error: %v", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&pokemon)
	if err != nil {
		return pokemon, fmt.Errorf("encountered error: %v", err)
	}

	return pokemon, nil
}
