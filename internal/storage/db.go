package storage

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB

func ConnectDb() {
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	pass := viper.GetString("DB_PASSWORD")
	name := viper.GetString("DB_NAME")
	host := viper.GetString("DB_HOST")
	dsn := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable", port, host, user, pass, name)

	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		panic("err")
	}
	err = Db.AutoMigrate(&User{}, &Account{}, &Record{}, &IncomeCategory{})
	if err != nil {
		logrus.Error(err.Error())
	}
	InitIncomeCategory()
}

type Model struct {
	ID        uint64    `gorm:"primary_key,autoincrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"Updated_at"`
	DeletedAt time.Time `json:"Deleted_at"`
}
