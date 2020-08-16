package model

type Room struct {
	ID       int64     `gorm:"column:room.id;primary_key;auto_increment"`
	Name     string    `gorm:"column:name"`
	Comments []Comment `gorm:"foreignKey:RoomID"`
}
