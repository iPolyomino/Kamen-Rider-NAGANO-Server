package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
)

type MockDB struct {
	ctrl  *gomock.Controller
	rwdb  *sql.DB
	rwm   sqlmock.Sqlmock
	rodb  *sql.DB
	rom   sqlmock.Sqlmock
	rwgdb *gorm.DB
	rogdb *gorm.DB
}

func (db MockDB) ExpectationsWereMet() error {
	err := db.rwm.ExpectationsWereMet()
	if err != nil {
		return err
	}
	err = db.rom.ExpectationsWereMet()
	if err != nil {
		return err
	}
	return nil
}

func (db MockDB) Close() {
	db.rogdb.Close()
	db.rwgdb.Close()
	db.rodb.Close()
	db.rwdb.Close()
	db.ctrl.Finish()
}
func GetMockDB(t *testing.T) (*MockDB, error) {
	ctrl := gomock.NewController(t)
	rwdb, rwm, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	rodb, rom, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	rwgdb, err := gorm.Open("mysql", rwdb)
	if err != nil {
		return nil, err
	}
	rogdb, err := gorm.Open("mysql", rodb)
	if err != nil {
		return nil, err
	}
	result := MockDB{
		ctrl:  ctrl,
		rwdb:  rwdb,
		rwm:   rwm,
		rodb:  rodb,
		rom:   rom,
		rwgdb: rwgdb,
		rogdb: rogdb,
	}
	return &result, err
}
