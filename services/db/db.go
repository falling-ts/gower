package db

import (
	"database/sql"
	"gorm.io/gorm"
	"gower/services"
	"time"
)

type DB struct {
	*gorm.DB
}

var configs services.Configs

// New 创建 DB
func New() *DB {
	return new(DB)
}

// Init 服务初始化
func (d *DB) Init(args ...any) {
	if len(args) == 0 {
		panic("数据库服务初始化参数不全.")
	}
	configs = args[0].(services.Configs)

	db, err := gorm.Open(driver(configs.Get("db.driver", "mysql").(string)), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: configs.Get("db.disableForeignKey", true).(bool),
	})
	if err != nil {
		panic(err)
	}

	d.DB = db

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(configs.Get("db.maxOpen", 100).(int))
	sqlDB.SetMaxIdleConns(configs.Get("db.maxIdleCount", 25).(int))
	sqlDB.SetConnMaxLifetime(configs.Get("db.maxLifeTime", "30m").(time.Duration))
	sqlDB.SetConnMaxIdleTime(configs.Get("db.maxIdleTime", "10m").(time.Duration))
}

// GormDB 获取 gorm DB
func (d *DB) GormDB() *gorm.DB {
	return d.DB
}

// SqlDB 获取 sql DB
func (d *DB) SqlDB() (*sql.DB, error) {
	return d.DB.DB()
}
