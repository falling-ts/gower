package models

import (
	"errors"
)

func init() {
	au := new(AdminUser)
	migrate(au)

	result := db.First(au, 1)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			au.Username = util.Ptr("admin").(*string)
			au.Password, _ = passwd.Hash(util.SHA256("admin"))
			au.Email = util.Ptr("admin@admin.com").(*string)
			au.Nickname = util.Ptr("Admin").(*string)

			db.Create(au)
		}
	}

}

type AdminUser struct {
	Model
	Username *string `gorm:"type:string;uniqueIndex;comment:用户名"`
	Password string  `gorm:"type:string;comment:密码"`
	Email    *string `gorm:"type:string;uniqueIndex;comment:邮箱"`
	Nickname *string `gorm:"type:string;default:'';comment:昵称"`
	Avatar   *string `gorm:"type:string;default:'';comment:头像"`
}

// From 从用户名获取数据
func (u *AdminUser) From(account string) error {
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
func (u *AdminUser) Login(aud ...string) (string, error) {
	return auth.Sign("admin_model", u.IDString(), aud)
}
