package storage

type Account struct {
	Id      int      `json:"id"`
	UserID  int      `json:"user_id" `
	Title   string   `json:"title"`
	Incomes []Record `json:"incomes"`
}
