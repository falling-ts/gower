package middlewares

import (
	"github.com/falling-ts/gower/app/middlewares"
	"github.com/falling-ts/gower/app/models"
	"github.com/falling-ts/gower/services"
)

var _ = Auth()

func Auth() services.Handler {
	return middlewares.Auth("admin-token", "Admin-Authorization", func(id string) (*models.Auth, error) {
		adminUser := new(models.AdminUser)
		result := db.First(adminUser, id)
		if result.Error != nil {
			return nil, trans.DBError(result.Error)
		}

		return &models.Auth{AdminUser: *adminUser}, nil
	})
}
