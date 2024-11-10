package middlewares

import (
	"github.com/falling-ts/gower/app"
	"github.com/falling-ts/gower/app/models"
	"github.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Menus() services.Handler {
	return func(c *gin.Context) {
		var menuModels []models.AdminMenu

		menusKey := "admin_menus"

		result := db.Preload("Children").Where("parent_id = ?", 0).Find(&menuModels)
		if result.Error != nil {
			excp.New(http.StatusBadRequest, result.Error).Handle(c)
			c.Abort()
			return
		}

		menus := make([]any, len(menuModels))
		for i, menu := range menuModels {
			tmp, err := menu.SetModel(&menu).Out(app.Rule{
				"id":   "ID",
				"icon": "Icon",
				"name": "Name",
				"path": "Path",
				"children": app.Rule{
					"id":   "ID",
					"icon": "Icon",
					"name": "Name",
					"path": "Path",
				},
			})
			if err != nil {
				excp.New(http.StatusBadRequest, result.Error).Handle(c)
				c.Abort()
				return
			}

			menus[i] = tmp
		}

		c.Set(menusKey, menus)
		c.Next()
	}
}
