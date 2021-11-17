package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to homepage from Bibek")
	fmt.Println("Endpoint Hit:homepage")
}

func handleRequests() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":1000", nil))
}

func main() {
	handleRequests()
}