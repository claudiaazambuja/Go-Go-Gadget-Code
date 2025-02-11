package services

import (
	"Go-Go-Gadget-Code/models"
	"Go-Go-Gadget-Code/repository"
)

func GetUsers() []models.User {
	return repository.GetAll()
}
