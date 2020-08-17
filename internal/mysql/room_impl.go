package mysql

import (
	"fmt"

	"github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/mysql/model"
	"github.com/jinzhu/gorm"
)

type DefaultRoom struct{}

func NewRoom() Room {
	return &DefaultRoom{}
}

func (d *DefaultRoom) Find(tx *gorm.DB, roomID int) (*model.Room, error) {
	var m model.Room
	query := tx.Table("room").
		Select("*").
		Joins("LEFT JOIN comment ON room.id = comment.room_id").
		Where("room.id = ?", roomID)
	res := query.Find(&m)
	if res.RecordNotFound() {
		return nil, fmt.Errorf("room_id=%d not found: %w", roomID, ErrNotFound)
	}
	if res.Error != nil {
		return nil, res.Error
	}

	rows, err := query.Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment model.Comment
		err := query.ScanRows(rows, &comment)
		if err != nil {
			return nil, err
		}
		m.Comments = append(m.Comments, comment)
	}

	return &m, nil
}
