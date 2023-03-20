package models

func init() {
	migrate(new(User))
}

type User struct {
	Model
	Username *string `gorm:"type:string;uniqueIndex;comment:用户名"`
	Password string  `gorm:"type:string;comment:密码"`
	Email    *string `gorm:"type:string;uniqueIndex;comment:邮箱"`
}

// Register 注册
func (u *User) Register() error {
	result := db.Create(u)
	return trans.DBError(result.Error)
}

func (u *User) First(username string) error {
	result := db.Where("username = ?", username).First(u)
	return trans.DBError(result.Error)
}
