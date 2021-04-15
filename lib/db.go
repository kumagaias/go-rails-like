package lib

import (
	"github.com/jszwec/csvutil"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"opb/config"
	"opb/models"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open(config.Get("db.database")),
		&gorm.Config{},
	)
	if err != nil {
		panic("DB error")
	}
	return db
}

func InitDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func SeedDB(db *gorm.DB) {
	var users []models.User
	b, _ := ioutil.ReadFile("./data/users.csv")
	csvutil.Unmarshal(b, &users)
	db.Create(&users)
}
