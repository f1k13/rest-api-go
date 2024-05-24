package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"rest_api_go/internal/storage"
	"strconv"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var creds storage.Account
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Неверный формат", http.StatusBadRequest)
		return
	}
	if err := storage.Db.Table("accounts").Create(&creds).Error; err != nil {
		http.Error(w, "Ошибка при создании счета", http.StatusInternalServerError)
		logrus.Error(err.Error())
		return
	}

	jsonRes, err := json.Marshal(creds)
	if err != nil {
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		logrus.Error(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonRes)
}

func GetAccOfUserId(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("id")
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		logrus.Error(err.Error())
		http.Error(w, "Неверный формат userId", http.StatusBadRequest)
		return
	}
	var account storage.Account

	if err := storage.Db.Where("user_id = ?", userId).First(&account).Error; err != nil {
		http.Error(w, "Счёт не найден", http.StatusNotFound)
		logrus.Error(err.Error())
		return
	}
	jsonRes, err := json.Marshal(account)

	if err != nil {
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		logrus.Error(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
