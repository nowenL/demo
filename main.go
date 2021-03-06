package main

import (
	"fmt"
	"net/http"
)

const (
	template = `<h1 style="text-align: center; font-size: 72px; margin-top: 200px;">%s</h1>`
)

func handler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf(template, "Hello cloud v1.0")
	fmt.Println("Response with: " + message)
	fmt.Fprintf(w, message)
}

func main() {

	fmt.Println()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
