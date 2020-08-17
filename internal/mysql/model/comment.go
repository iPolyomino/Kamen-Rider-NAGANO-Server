package model

type Comment struct {
	ID       int64  `gorm:"column:comment.id;primary_key;auto_increment"`
	Sender   string `gorm:"column:sender"`
	Text     string `gorm:"column:text"`
	ImageURL string `gorm:"column:image_url"`
	RoomID   int64
}
