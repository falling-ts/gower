package models

import (
	"gorm.io/gorm"
	"gower/app"
	"strconv"
)

type Model struct {
	app.ModelHandle `gorm:"-"`
	gorm.Model
}

var (
	db    = app.DB()
	trans = app.Translate()
	token = app.Token()
)

func migrate(args ...any) {
	if err := db.AutoMigrate(args...); err != nil {
		panic(err)
	}
}

// IDString 获取 ID 字符串
func (m *Model) IDString() string {
	return strconv.FormatUint(uint64(m.ID), 10)
}
