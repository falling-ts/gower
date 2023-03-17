package providers

import (
	"gower/services"
	"gower/services/db"
)

var _ services.DBService = (*db.Service)(nil)

func init() {
	ss.DB = db.New()
}
