package controllers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
