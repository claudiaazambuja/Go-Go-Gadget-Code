package handlers

import (
	"encoding/json"
	"net/http"

	"Go-Go-Gadget-Code/services"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Verifica se o método é GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os usuários do serviço
	users := services.GetUsers()

	// Codifica os usuários como JSON e envia a resposta
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Erro ao gerar JSON", http.StatusInternalServerError)
	}
}
