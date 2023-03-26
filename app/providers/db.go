package providers

import (
	"github.com/falling-ts/gower/services"
	"github.com/falling-ts/gower/services/db"
)

var _ services.DBService = (*db.Service)(nil)

func init() {
	P.Register("db", func() (Depends, Resolve) {
		return Depends{"config", "logger"}, func(ss ...services.Service) services.Service {
			return db.New().Init(ss...)
		}
	})
}
