package models

func init() {
	migrate(new(AdminMenu))

	var count int64
	db.Model(&AdminMenu{}).Count(&count)
	if count == 0 {
		menus := []*AdminMenu{
			{
				Icon:     util.Ptr("line-md--home").(*string),
				Name:     "主页",
				Path:     util.Ptr("/admin").(*string),
				Order:    1,
				ParentID: util.Ptr(uint(0)).(*uint),
			},
			{
				Icon:     util.Ptr("line-md--cog-loop").(*string),
				Name:     "系统设置",
				Order:    2,
				ParentID: util.Ptr(uint(0)).(*uint),
				Children: []AdminMenu{
					{
						Icon:  util.Ptr("line-md--person-add").(*string),
						Name:  "员工管理",
						Path:  util.Ptr("/admin/user").(*string),
						Order: 3,
					},
					{
						Icon:  util.Ptr("line-md--person-filled").(*string),
						Name:  "角色管理",
						Path:  util.Ptr("/admin/role").(*string),
						Order: 4,
					},
					{
						Icon:  util.Ptr("line-md--cancel-twotone").(*string),
						Name:  "权限管理",
						Path:  util.Ptr("/admin/permission").(*string),
						Order: 5,
					},
					{
						Icon:  util.Ptr("line-md--align-left").(*string),
						Name:  "菜单管理",
						Path:  util.Ptr("/admin/menu").(*string),
						Order: 6,
					},
				},
			},
		}

		db.Create(menus)
	}
}

type AdminMenu struct {
	Model

	Icon  *string `gorm:"type:string;comment:图标"`
	Name  string  `gorm:"type:string;not null;comment:菜单名称"`
	Path  *string `gorm:"type:string;comment:菜单路由"`
	Order uint    `gorm:"type:uint;default:0;not null;commit:排序"`

	ParentID *uint
	Parent   *AdminMenu `gorm:"foreignKey:ParentID"`

	Children []AdminMenu `gorm:"foreignKey:ParentID"`

	Roles []*AdminRole `gorm:"many2many:admin_menu_roles;joinForeignKey:menu_id;joinReferences:role_id"`
}
