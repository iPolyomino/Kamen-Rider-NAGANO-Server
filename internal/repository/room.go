package repository

import "github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/mysql/model"

type Room interface {
	GetRoom(roomID int) (*model.Room, error)
}
