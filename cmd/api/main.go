package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type Pokemon struct {
	Name string `json:"name"`
}

func getPokemonFromEnv() []Pokemon {
	var pokemons []Pokemon
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if strings.HasPrefix(pair[0], "POKEMON_") {
			pokemons = append(pokemons, Pokemon{Name: pair[1]})
		}
	}
	return pokemons
}

var inMemoryPokemons = []Pokemon{
	{Name: "Bulbasaur"},
	{Name: "Charmander"},
	{Name: "Squirtle"},
}

func getAllPokemons() []Pokemon {
	return append(inMemoryPokemons, getPokemonFromEnv()...)
}

func listPokemonsHandler(w http.ResponseWriter, r *http.Request) {
	pokemons := getAllPokemons()
	json.NewEncoder(w).Encode(pokemons)
}

func main() {
	http.HandleFunc("/", listPokemonsHandler)
	http.ListenAndServe(":8080", nil)
}
