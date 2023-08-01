package image_entity

type Image struct {
	ID  uint8
	URL string `gorm:"type:varchar(3000)"`
}
