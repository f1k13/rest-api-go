package storage

type IncomeCategory struct {
	Id    int    `json:"Id"`
	Title string `json:"Title"`
	Color string `json:"Color"`
	Icon  string `json:"Icon"`
}

func InitIncomeCategory() {
	var count int64
	Db.Model(&IncomeCategory{}).Count(&count)
	if count > 0 {
		return
	}
	income := []IncomeCategory{
		{Title: "Зарплата", Color: "#FFD700", Icon: "Salary"},
		{Title: "Продажа", Color: "#00FF00", Icon: "Sales"},
		{Title: "Аренда", Color: "#000FF", Icon: "Lease"},
	}
	for _, category := range income {
		Db.Create(&category)
	}
}
