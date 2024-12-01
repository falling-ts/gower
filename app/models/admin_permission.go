package models

func init() {
	migrate(new(AdminPermission))
}

type AdminPermission struct {
	Model

	Name   *string `gorm:"type:string;comment:权限名称"`
	Flag   string  `gorm:"type:string;not null;comment:权限标识"`
	Method string  `gorm:"type:string;default:'ANY';comment:请求方法"`
	Path   string  `gorm:"type:string;not null;comment:请求路由"`
	Order  uint    `gorm:"type:uint;default:0;not null;commit:排序"`

	ParentID *uint
	Parent   *AdminPermission `gorm:"foreignKey:ParentID"`

	Children []AdminPermission `gorm:"foreignKey:ParentID"`

	Roles []*AdminRole `gorm:"many2many:admin_permission_roles;joinForeignKey:permission_id;joinReferences:role_id"`
}
