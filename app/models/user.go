package models

import "errors"

func init() {
	migrate(new(User))
}

type User struct {
	Model
	Username *string `gorm:"type:string;uniqueIndex;comment:用户名"`
	Password string  `gorm:"type:string;comment:密码"`
	Email    *string `gorm:"type:string;uniqueIndex;comment:邮箱"`
	Nickname *string `gorm:"type:string;default:'';comment:昵称"`
	Avatar   *string `gorm:"type:string;default:'';comment:头像"`
}

// Register 注册
func (u *User) Register() error {
	result := db.Create(u)
	return trans.DBError(result.Error)
}

// From 从用户名获取数据
func (u *User) From(account string) error {
	result := db.Where("username = ?", account).First(u)
	if result.Error != nil {
		result = db.Where("email = ?", account).First(u)
	}
	if result.Error != nil && result.Error.Error() == "record not found" {
		return errors.New("账号未注册")
	}
	return trans.DBError(result.Error)
}

// Login 登录
func (u *User) Login(aud ...string) (string, error) {
	return auth.Sign("user_model", u.IDString(), aud)
}
