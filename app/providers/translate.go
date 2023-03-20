package providers

import (
	"gower/services"
	"gower/services/translate"
	"gower/trans"
)

var _ services.TranslateService = (*translate.Service)(nil)

func init() {
	ss.Translate = translate.Mount(trans.All)
}
