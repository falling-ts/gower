package db

import (
	"fmt"
	"time"

	rawMysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func driver(driver string) gorm.Dialector {
	switch driver {
	case "mysql":
		return getMysqlDriver()
	default:
		return nil
	}
}

func getMysqlDriver() gorm.Dialector {
	loc, ok := map[string]*time.Location{
		"sys":     time.Local,
		"utc":     time.UTC,
		"beijing": time.FixedZone("Asia/Shanghai", int((8 * time.Hour).Seconds())),
	}[config.Get("db.timezone", "sys").(string)]
	if !ok {
		loc = time.Local
	}
	dsnConfig := &rawMysql.Config{
		User:   config.Get("db.user", "root").(string),
		Passwd: config.Get("db.passwd").(string),
		Net:    config.Get("db.net", "tcp").(string),
		Addr: fmt.Sprintf("%s:%d",
			config.Get("db.host", "localhost").(string),
			config.Get("db.port", 3306).(int)),
		DBName:               config.Get("db.name", "gower").(string),
		AllowNativePasswords: config.Get("db.mysql.allowNativePasswords", true).(bool),
		Loc:                  loc,
	}
	return mysql.New(mysql.Config{
		DSNConfig:                 dsnConfig,
		DSN:                       dsnConfig.FormatDSN(),
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	})
}
