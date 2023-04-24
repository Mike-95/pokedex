package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreName := args[0]

	locationsArea, err := cfg.pokeapiClient.GetLocationArea(locationAreName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:\n", locationsArea.Name)
	for _, pokemon := range locationsArea.PokemonEncounters {
		fmt.Printf("-%s\n", pokemon.Pokemon.Name)
	}

	return nil
}
