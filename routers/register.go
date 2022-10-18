package routers

import (
	"encoding/json"
	"net/http"
	"twittor/db"
	"twittor/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Erro do dados recebidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "O email é obrigatório", 400)
		return
	}

	if len(t.Email) < 6 {
		http.Error(w, "O email tem q ter no mínimo 6 caracteres", 400)
		return
	}

	_, findEmail, _ := db.CheckUser(t.Email)
	if findEmail {
		http.Error(w, "Já existe usuário com esse email", 400)
		return
	}

	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "Ocorreu um erro ao realizar o registro de usuário "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "erro no registro de usuário ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
