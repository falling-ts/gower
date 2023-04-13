package models

func init() {
	migrate(new(AdminPermission))
}

type AdminPermission struct {
	Model

	Name   *string `gorm:"type:string;comment:权限名称"`
	Method string  `gorm:"type:string;default:'ANY';comment:请求方法"`
	Path   string  `gorm:"type:string;not null;comment:请求路由"`
}
