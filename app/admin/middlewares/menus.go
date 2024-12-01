package middlewares

import (
	"gitee.com/falling-ts/gower/app"
	"gitee.com/falling-ts/gower/app/models"
	"gitee.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Menus() services.Handler {
	return func(c *gin.Context) {
		var menuModels []models.AdminMenu

		result := db.Preload("Children").Where("parent_id = ?", 0).Find(&menuModels)
		if result.Error != nil {
			exc.New(http.StatusBadRequest, result.Error).Handle(c)
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
				exc.New(http.StatusBadRequest, result.Error).Handle(c)
				c.Abort()
				return
			}

			menus[i] = tmp
		}

		c.Set("admin_menus", menus)
		c.Next()
	}
}
