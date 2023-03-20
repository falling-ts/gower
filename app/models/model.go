package models

import (
	"gorm.io/gorm"
	"gower/app"
)

type Model struct {
	app.ModelHandle `gorm:"-"`
	gorm.Model
}

var (
	db    = app.DB()
	trans = app.Translate()
)

func migrate(args ...any) {
	if err := db.AutoMigrate(args...); err != nil {
		panic(err)
	}
}
