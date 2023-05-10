package image_entity

type Image struct {
	ID  uint
	URL string `gorm:"type:varchar(300)"`
}
