package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		if strings.Contains(userAgent, "Firefox") {
			fmt.Fprint(w, "Testing is allowed!\n")
		} else {
			fmt.Fprint(w, "Welcome!\n")
		}
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
