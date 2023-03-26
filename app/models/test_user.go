package models

func init() {
	migrate(new(TestUser))
}

type TestUser struct {
	Model
	Username   *string        `gorm:"type:string;uniqueIndex;comment:用户名"`
	Password   string         `gorm:"type:string;comment:密码"`
	Email      *string        `gorm:"type:string;uniqueIndex;comment:邮箱"`
	UserInfo   TestUserInfo   `gorm:"foreignKey:UserID"`
	Categories []TestCategory `gorm:"foreignKey:UserID"`
	Articles   []TestArticle  `gorm:"foreignKey:UserID"`
	Comments   []TestComment  `gorm:"foreignKey:UserID"`
}
