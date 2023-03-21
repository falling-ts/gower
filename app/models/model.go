package models

import (
	"gorm.io/gorm"
	"gower/app"
	"strconv"
	"time"
)

type Model struct {
	app.Model `gorm:"-"`
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
