package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome!!!!\n")
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
