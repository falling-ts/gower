package models

func init() {
	migrate(new(User))
}

type User struct {
	Model
	Username string `gorm:"type:string;uniqueIndex;comment:用户名"`
}
