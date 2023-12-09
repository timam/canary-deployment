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
	printRoleLabel()

	pokemons := getAllPokemons()
	response := PokemonsResponse{Pokemons: pokemons}
	json.NewEncoder(w).Encode(response)
}

func printRoleLabel() {
	config, _ := rest.InClusterConfig()
	clientset, _ := kubernetes.NewForConfig(config)

	podName := os.Getenv("POD_NAME")
	namespace := os.Getenv("NAMESPACE")

	fmt.Println(podName)
	fmt.Println(namespace)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pod, _ := clientset.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
	labels := pod.GetLabels()

	if role, ok := labels["role"]; ok {
		fmt.Println(role)
	}

	cancel() // Cancel the context as soon as we're done with it
}

func main() {
	http.HandleFunc("/", listPokemonsHandler)
	fmt.Println("Pokemons are out")
	http.ListenAndServe(":8080", nil)
}
