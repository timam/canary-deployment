package main

import (
	"encoding/json"
	"fmt"
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

type PokemonsResponse struct {
	Pokemons []Pokemon `json:"pokemons"`
}

func listPokemonsHandler(w http.ResponseWriter, r *http.Request) {
	podDetails()

	pokemons := getAllPokemons()
	response := PokemonsResponse{Pokemons: pokemons}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", listPokemonsHandler)
	http.HandleFunc("/healthz", healthz)
	fmt.Println("Pokemons are out")
	http.ListenAndServe(":8080", nil)
}
