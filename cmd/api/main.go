package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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
	printAnnotations()

	pokemons := getAllPokemons()
	response := PokemonsResponse{Pokemons: pokemons}
	json.NewEncoder(w).Encode(response)
}

func printAnnotations() {
	config, _ := rest.InClusterConfig()
	clientset, _ := kubernetes.NewForConfig(config)

	podName := os.Getenv("POD_NAME")
	namespace := os.Getenv("NAMESPACE")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pod, _ := clientset.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
	annotations := pod.GetAnnotations()

	for key, value := range annotations {
		fmt.Printf("%s: %s\n", key, value)
	}

	cancel() // Cancel the context as soon as we're done with it
}

func main() {
	http.HandleFunc("/", listPokemonsHandler)
	http.ListenAndServe(":8080", nil)
}
