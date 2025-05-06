package funcs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLocations(url string) (LocationAreas, error) {

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

func GetPokemon(url string) (LocationAreaEncounters, error) {
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
