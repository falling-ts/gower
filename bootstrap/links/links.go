package links

import (
	"gower/app"
	"gower/app/exceptions"
	"gower/configs"
)

func init() {
	new(configs.Configs).Link(app.Gower.Config())
	new(exceptions.Exceptions).Link(app.Gower.Exception())
}
