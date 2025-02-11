package repository

import "Go-Go-Gadget-Code/models"

var users = []models.User{
	{ID: 1, Name: "Alice", Email: "alice@email.com"},
	{ID: 2, Name: "Bob", Email: "bob@email.com"},
}

// GetAll retorna todos os usuários
func GetAll() []models.User {
	return users
}
