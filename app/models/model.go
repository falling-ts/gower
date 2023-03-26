package models

import (
	"github.com/falling-ts/gower/app"
	"gorm.io/gorm"
	"strconv"
)

type Model struct {
	app.ModelHandle `gorm:"-"`
	gorm.Model
}

var (
	db    = app.DB()
	trans = app.Translate()
	auth  = app.Auth()
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
