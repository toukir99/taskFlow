package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"task-gateway/db"
	"task-gateway/web/model"
	"task-gateway/web/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := utils.Validate(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := user.HashPassword(); err != nil {
		slog.Error("Error hashing password", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	userId, err := db.GetAuthRepo().CreateUser(&user)
	if err != nil {
		slog.Error("Error creating user", "err", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err := SendEmail(user.Email); err != nil {
		utils.SendError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.SendData(w, userId)
}
