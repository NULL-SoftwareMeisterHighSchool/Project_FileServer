package image_entity

type Image struct {
	ID  uint   `gorm:"autoIncrement"`
	URL string `gorm:"type:varchar(3000)"`
}
