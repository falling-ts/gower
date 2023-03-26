package models

func init() {
	migrate(new(TestComment))
}

type TestComment struct {
	Model
	Content   *string `gorm:"type:text;comment:评论内容"`
	UserID    *uint
	User      *TestUser `gorm:"foreignKey:UserID"`
	ArticleID *uint
	Article   *TestArticle `gorm:"foreignKey:ArticleID"`
	ParentID  *uint
	Parent    *TestComment  `gorm:"foreignKey:ParentID"`
	Children  []TestComment `gorm:"foreignKey:ParentID"`
}
