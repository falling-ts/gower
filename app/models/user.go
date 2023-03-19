package models

func init() {
	migrate(new(User))
}

type User struct {
	Model
	Username string `gorm:"type:string;uniqueIndex;comment:用户名"`
	Password string `gorm:"type:string;comment:密码"`
	Email    string `gorm:"type:string;uniqueIndex;comment:邮箱"`
}

// Register 注册
func (u *User) Register() error {
	result := db.Create(u)
	return result.Error
}
