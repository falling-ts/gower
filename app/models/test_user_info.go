package models

func init() {
	migrate(new(TestUserInfo))
}

type TestUserInfo struct {
	Model
	Nickname *string `gorm:"type:string;default:'';comment:昵称"`
	Avatar   *string `gorm:"type:string;default:'';comment:头像"`
	UserID   *uint
	User     *TestUser `gorm:"foreignKey:UserID"`
}
