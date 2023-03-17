package services

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DBService interface {
	Service

	AutoMigrate(dst ...interface{}) error
	Association(column string) *gorm.Association
	Attrs(attrs ...interface{}) (tx *gorm.DB)
	Assign(attrs ...interface{}) (tx *gorm.DB)
	AddError(err error) error

	Begin(opts ...*sql.TxOptions) *gorm.DB

	Create(value interface{}) (tx *gorm.DB)
	CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB)
	Count(count *int64) (tx *gorm.DB)
	Connection(fc func(tx *gorm.DB) error) (err error)
	Commit() *gorm.DB
	Clauses(conds ...clause.Expression) (tx *gorm.DB)

	Distinct(args ...interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	Debug() (tx *gorm.DB)
	SqlDB() (*sql.DB, error)

	Exec(sql string, values ...interface{}) (tx *gorm.DB)

	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB
	FirstOrInit(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	FirstOrCreate(dest interface{}, conds ...interface{}) (tx *gorm.DB)

	Group(name string) (tx *gorm.DB)
	Get(key string) (interface{}, bool)
	GormDB() *gorm.DB

	Having(query interface{}, args ...interface{}) (tx *gorm.DB)

	InnerJoins(query string, args ...interface{}) (tx *gorm.DB)
	InstanceSet(key string, value interface{}) *gorm.DB
	InstanceGet(key string) (interface{}, bool)

	Joins(query string, args ...interface{}) (tx *gorm.DB)

	Last(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Limit(limit int) (tx *gorm.DB)

	Migrator() gorm.Migrator
	Model(value interface{}) (tx *gorm.DB)

	Not(query interface{}, args ...interface{}) (tx *gorm.DB)

	Omit(columns ...string) (tx *gorm.DB)
	Or(query interface{}, args ...interface{}) (tx *gorm.DB)
	Order(value interface{}) (tx *gorm.DB)
	Offset(offset int) (tx *gorm.DB)

	Pluck(column string, dest interface{}) (tx *gorm.DB)
	Preload(query string, args ...interface{}) (tx *gorm.DB)

	Raw(sql string, values ...interface{}) (tx *gorm.DB)
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	Rollback() *gorm.DB
	RollbackTo(name string) *gorm.DB

	Save(value interface{}) (tx *gorm.DB)
	Scan(dest interface{}) (tx *gorm.DB)
	ScanRows(rows *sql.Rows, dest interface{}) error
	SavePoint(name string) *gorm.DB
	Select(query interface{}, args ...interface{}) (tx *gorm.DB)
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) (tx *gorm.DB)
	Session(config *gorm.Session) *gorm.DB
	Set(key string, value interface{}) *gorm.DB
	SetupJoinTable(model interface{}, field string, joinTable interface{}) error

	Table(name string, args ...interface{}) (tx *gorm.DB)
	Take(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error)
	ToSQL(queryFn func(tx *gorm.DB) *gorm.DB) string

	Use(plugin gorm.Plugin) error
	Unscoped() (tx *gorm.DB)
	Update(column string, value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	UpdateColumn(column string, value interface{}) (tx *gorm.DB)
	UpdateColumns(values interface{}) (tx *gorm.DB)

	WithContext(ctx context.Context) *gorm.DB
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
}
