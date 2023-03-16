package providers

import (
	"gower/services"
	"gower/services/db"
)

var _ services.DB = (*db.DB)(nil)

func init() {
	s.DB = db.New()
}
