package mysql

import (
	"errors"
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/mysql/model"
)

func NewRoomRows(ms ...model.Room) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{
		"room.id",
		"name",
		"comment.id",
		"sender",
		"text",
		"image_url",
		"room_id",
	})

	for _, m := range ms {
		for _, mc := range m.Comments {
			rows.AddRow(m.ID, m.Name, mc.ID, mc.Sender, mc.Text, mc.ImageURL, mc.RoomID)
		}
	}

	return rows
}

func TestDefaultRoom_Find(t *testing.T) {
	query := "SELECT * FROM `room` LEFT JOIN comment ON room.id = comment.room_id WHERE (room.id = ?)"

	value := model.Room{
		ID:   1,
		Name: "hoge",
		Comments: []model.Comment{
			{
				ID:       2,
				Sender:   "fuga",
				Text:     "fugafuga",
				ImageURL: "https://fuga.fuga",
			},
			{
				ID:       3,
				Sender:   "piyo",
				Text:     "piyopiyo",
				ImageURL: "https://piyo.piyo",
			},
		},
	}

	testCases := TestCases{
		{
			name: "OK",
			run: func(t *testing.T, tx *gorm.DB) error {
				res, err := NewRoom().Find(tx, 1)

				if diff := cmp.Diff(&value, res); len(diff) != 0 {
					t.Errorf("diff=%s", diff)
				}
				return err
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(1).
					WillReturnRows(NewRoomRows(value))
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(1).
					WillReturnRows(NewRoomRows(value))
				mock.ExpectCommit()
			},
			check: Succeeded,
		},
		{
			name: "レコードが存在しない",
			run: func(t *testing.T, tx *gorm.DB) error {
				_, err := NewRoom().Find(tx, 1)
				return err
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(1).
					WillReturnRows(NewRoomRows())
				mock.ExpectRollback()
			},
			check: Failed(ErrNotFound),
		},
		{
			name: "query.Findでなんらかのエラーで失敗した場合",
			run: func(t *testing.T, tx *gorm.DB) error {
				_, err := NewRoom().Find(tx, 1)
				return err
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(1).
					WillReturnError(errors.New(""))
				mock.ExpectRollback()
			},
			check: FailedAny,
		},
		{
			name: "query.Rowsでなんらかのエラーで失敗した場合",
			run: func(t *testing.T, tx *gorm.DB) error {
				_, err := NewRoom().Find(tx, 1)
				return err
			},
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(1).
					WillReturnRows(NewRoomRows(value))
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(1).
					WillReturnError(errors.New(""))
				mock.ExpectRollback()
			},
			check: FailedAny,
		},
	}
	testCases.Run(t)
}
