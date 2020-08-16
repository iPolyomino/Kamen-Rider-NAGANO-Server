package mysql

import (
	"github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/mysql/model"
	"github.com/jinzhu/gorm"
)

type Room interface {
	Find(tx *gorm.DB, roomID int) (*model.Room, error)
}
