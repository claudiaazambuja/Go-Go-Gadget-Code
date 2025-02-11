package main

import (
	"fmt"
	"log"
	"net/http"

	"Go-Go-Gadget-Code/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.GetUsers)

	port := ":8080"
	fmt.Println("🚀 Servidor rodando na porta", port)
	log.Fatal(http.ListenAndServe(port, nil))
}