package models

import (
	"gorm.io/gorm"
	"gower/app"
)

var db = app.GormDB()

type Model struct {
	gorm.Model
}

func migrate(args ...any) {
	if err := db.AutoMigrate(args...); err != nil {
		panic(err)
	}
}
