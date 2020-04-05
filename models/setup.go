package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("sqlite3", "tasks.db")

	if err != nil {
		panic("Não foi possível conectar com banco de dados!")
	}

	db.AutoMigrate(&Task{})

	return db
}
