package entity

type Tag struct {
	Id   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" validate:"required"`
}

func NewTag(tagID uint, tagName string) *Tag {
	return &Tag{
		Id:   tagID,
		Name: tagName,
	}
}
