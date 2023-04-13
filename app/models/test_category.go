package models

func init() {
	migrate(new(TestCategory))
}

type TestCategory struct {
	Model
	Name  string `gorm:"type:string;default:'';not null;comment:类别名称"`
	Order uint   `gorm:"type:uint;default:0;not null;commit:排序"`

	ParentID *uint
	Parent   *TestCategory `gorm:"foreignKey:ParentID"`

	Children []TestCategory `gorm:"foreignKey:ParentID"`

	UserID *uint
	User   *TestUser `gorm:"foreignKey:UserID"`

	Articles []TestArticle `gorm:"foreignKey:CategoryID"`
}
