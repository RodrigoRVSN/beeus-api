package entity

type Tag struct {
	Id   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" validate:"required"`
}
