package models

func init() {
	migrate(new(TestArticle))
}

type TestArticle struct {
	Model
	Title      string  `gorm:"type:string;default:'';not null;commit:标题"`
	Content    *string `gorm:"type:text;commit:内容"`
	CategoryID *uint
	Category   *TestCategory `gorm:"foreignKey:CategoryID"`
	UserID     *uint
	User       *TestUser     `gorm:"foreignKey:UserID"`
	Comments   []TestComment `gorm:"foreignKey:ArticleID"`
}
