package handlers

import (
	"encoding/json"
	"net/http"
	"rest_api_go/internal/storage"
	"rest_api_go/middlewares"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var isUser storage.User

	if err := json.NewDecoder(r.Body).Decode(&isUser); err != nil {
		http.Error(w, "Невозможно декодировать JSON", http.StatusBadRequest)
	}
	var user storage.User

	query := storage.Db.Table("users")

	query.Where("username = ?", isUser.Username)

	query.Find(&user)

	if user.Username == "" {
		http.Error(w, "Пользователь не найден", http.StatusBadRequest)
		return
	}

	token, err := middlewares.GenerateToken(user.Id)
	if err != nil {
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	res := map[string]string{"token": token}

	jsonRes, err := json.Marshal(res)

	if err != nil {
		http.Error(w, "Ошибка", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var creds storage.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Невозможно декодировать JSON", http.StatusBadRequest)
		return
	}

	var existUser storage.User
	query := storage.Db.Table("users")
	err := query.Where("username = ?", creds.Username).First(&existUser).Error

	if existUser.Username != "" {
		http.Error(w, "Пользователь с таким именем уже существует", http.StatusBadRequest)
		return
	}

	if err := storage.Db.Table("users").Create(&creds).Error; err != nil {
		http.Error(w, "Ошибка при создании пользователя", http.StatusBadRequest)
		return
	}

	token, err := middlewares.GenerateToken(creds.Id)
	if err != nil {
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	res := map[string]string{"token": token}
	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonRes)
}
