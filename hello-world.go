package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, R *http.Request) {
	fmt.Fprintln(w, "Saikat is a bad boy")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
