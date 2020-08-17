package repository

import (
	"github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/mysql"
	"github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/mysql/model"
	"github.com/jinzhu/gorm"
)

type DefaultRoom struct {
	rw, ro *gorm.DB
	mysql.Room
}

func NewRoom(rw, ro *gorm.DB) Room {
	return &DefaultRoom{
		rw:   rw,
		ro:   ro,
		Room: mysql.NewRoom(),
	}
}

func (r *DefaultRoom) GetRoom(roomID int) (*model.Room, error) {
	var room *model.Room
	err := r.ro.Transaction(func(tx *gorm.DB) error {
		rm, err := r.Find(tx, roomID)
		if err != nil {
			return err
		}
		room = rm
		return nil
	})

	if err != nil {
		return nil, err
	}

	return room, nil
}
