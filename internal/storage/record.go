package storage

type Record struct {
	Id         int    `json:"id"`
	AccountID  int    `json:"account_id"`
	CategoryID int    `json:"category_id"`
	Title      string `json:"title"`
	Type       string `json:"type"`
}
