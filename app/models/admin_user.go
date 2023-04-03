package models

func init() {
	migrate(new(AdminUser))
}

type AdminUser struct {
	Model
	Username *string `gorm:"type:string;uniqueIndex;comment:用户名"`
	Password string  `gorm:"type:string;comment:密码"`
	Email    *string `gorm:"type:string;uniqueIndex;comment:邮箱"`
	Nickname *string `gorm:"type:string;default:'';comment:昵称"`
	Avatar   *string `gorm:"type:string;default:'';comment:头像"`
}
