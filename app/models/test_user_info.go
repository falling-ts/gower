package models

func init() {
	migrate(new(TestUserInfo))
}

type TestUserInfo struct {
	Model
	Nickname *string `gorm:"type:string;default:'';comment:昵称"`
	Avatar   *string `gorm:"type:string;default:'';comment:头像"`
	UserId   uint    `gorm:"type:unit;default:0;not null"`
	User     *TestUser
}
