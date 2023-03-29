package providers

import (
	"github.com/falling-ts/gower/services"
	"github.com/falling-ts/gower/services/upload"
)

var _ services.UploadService = (*upload.Service)(nil)

func init() {
	P.Register("upload", Depends{"config", "util"}, func(ss ...services.Service) services.Service {
		return upload.New().Init(ss...)
	})
}
