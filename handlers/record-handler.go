package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"rest_api_go/internal/storage"
	"strconv"
)

type RecordRes struct {
	ID        int                    `json:"id"`
	AccountID int                    `json:"accountID"`
	Category  storage.IncomeCategory `json:"category"`
	Title     string                 `json:"title"`
	Type      string                 `json:"type"`
}

func CreateRec(w http.ResponseWriter, r *http.Request) {
	var creds storage.Record

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Ошибка при декодировании", http.StatusBadRequest)
		logrus.Error(err.Error())
		return
	}
	if creds.Type != "income" && creds.Type != "expence" {
		http.Error(w, "Тип может быть только income || expence", http.StatusBadRequest)
		return
	}
	if err := storage.Db.Table("records").Create(&creds).Error; err != nil {
		logrus.Error(err.Error())
		http.Error(w, "Ошибка при создании дохода", http.StatusInternalServerError)
		return
	}
	jsonRes, err := json.Marshal(creds)
	if err != nil {
		logrus.Error(err.Error())
		http.Error(w, "Ошибка", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
func GetCat(w http.ResponseWriter, r *http.Request) {
	var inc []storage.IncomeCategory

	if err := storage.Db.Find(&inc).Error; err != nil {
		logrus.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inc)
}

func GetCreatingRec(w http.ResponseWriter, r *http.Request) {
	accIdStr := r.URL.Query().Get("id")
	accId, err := strconv.Atoi(accIdStr)
	if err != nil {
		logrus.Error(err.Error())
		http.Error(w, "Неверный формат accId", http.StatusBadRequest)
		return
	}
	var incomes []storage.Record
	if err := storage.Db.Where("account_id = ?", accId).Find(&incomes).Error; err != nil {
		http.Error(w, "Доход не найден", http.StatusNotFound)
		logrus.Error(err.Error())
		return
	}
	var incRes []RecordRes
	for _, income := range incomes {
		var cat storage.IncomeCategory
		if err := storage.Db.First(&cat, income.CategoryID).Error; err != nil {
			http.Error(w, "Категория не найдена", http.StatusNotFound)
			logrus.Error(err.Error())
			return
		}
		incomeResponse := RecordRes{ID: income.Id, AccountID: income.AccountID, Category: cat, Title: income.Title, Type: income.Type}
		incRes = append(incRes, incomeResponse)
	}
	jsonRes, err := json.Marshal(incRes)

	if err != nil {
		http.Error(w, "Ошибка", http.StatusInternalServerError)
		logrus.Error(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

//func GetByType(w http.ResponseWriter, r *http.Request) {
//	typeRec := r.URL.Query().Get("type")
//	var record []storage.Record
//}
