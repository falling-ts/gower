package models

type TestCategory struct {
	Model
	
	Name string `gorm:"type:string;default:'';not null;comment:类别名称"`
}
