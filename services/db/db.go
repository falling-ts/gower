package db

import (
	"database/sql"
	"gorm.io/gorm"
	"gower/services"
	"time"
)

type Service struct {
	*gorm.DB
}

var config services.Config

// New 创建 DB
func New() *Service {
	return new(Service)
}

// Init 服务初始化
func (s *Service) Init(args ...any) {
	config = args[0].(services.Config)

	db, err := gorm.Open(driver(config.Get("db.driver", "mysql").(string)), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: config.Get("db.disableForeignKey", true).(bool),
	})
	if err != nil {
		panic(err)
	}

	s.DB = db

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(config.Get("db.maxOpen", 100).(int))
	sqlDB.SetMaxIdleConns(config.Get("db.maxIdleCount", 25).(int))
	sqlDB.SetConnMaxLifetime(config.Get("db.maxLifeTime", "30m").(time.Duration))
	sqlDB.SetConnMaxIdleTime(config.Get("db.maxIdleTime", "10m").(time.Duration))
}

// GormDB 获取 gorm DB
func (s *Service) GormDB() *gorm.DB {
	return s.DB
}

// SqlDB 获取 sql DB
func (s *Service) SqlDB() (*sql.DB, error) {
	return s.DB.DB()
}
