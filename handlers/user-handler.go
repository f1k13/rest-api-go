package handlers

import (
	"encoding/json"
	"net/http"
	"rest_api_go/internal/storage"
	"rest_api_go/middlewares"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Отсуствует токен", http.StatusBadRequest)
		return
	}
	userID, err := middlewares.ValidateToken(token)

	if err != nil {
		http.Error(w, "Невалидный токен", http.StatusUnauthorized)
		return
	}
	var user storage.User
	if err := storage.Db.Table("users").Where("id = ?", userID).First(&user).Error; err != nil {
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		return
	}
	jsonRes, err := json.Marshal(user)

	if err != nil {
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
