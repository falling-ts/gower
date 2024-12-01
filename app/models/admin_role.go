package models

func init() {
	migrate(new(AdminRole))
}

type AdminRole struct {
	Model

	Name *string `gorm:"type:string;comment:角色名称"`
	Flag string  `gorm:"type:string;not null;comment:角色标记"`

	Users []*AdminUser `gorm:"many2many:admin_role_users;joinForeignKey:role_id;joinReferences:user_id"`

	Menus []*AdminMenu `gorm:"many2many:admin_menu_roles;joinForeignKey:role_id;joinReferences:menu_id"`

	Permissions []*AdminPermission `gorm:"many2many:admin_permission_roles;joinForeignKey:role_id;joinReferences:permission_id"`
}
