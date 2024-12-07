package models

func init() {
	migrate(new({{.UpCamel}}))
}

type {{.UpCamel}} struct {
	Model

	// Name *string `gorm:"type:string;default:'';comment:名称"`
}
