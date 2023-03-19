package providers

import (
	"gower/services"
	"gower/services/passwd"
)

var _ services.PasswdService = (*passwd.Service)(nil)

func init() {
	ss.Passwd = new(passwd.Service)
}
