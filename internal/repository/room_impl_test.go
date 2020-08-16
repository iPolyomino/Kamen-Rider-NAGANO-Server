package repository

import (
	"errors"
	"testing"

	"github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/mysql"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/mysql/model"

	"github.com/DATA-DOG/go-sqlmock"
	mock_mysql "github.com/iPolyomino/Kamen-Rider-NAGANO-Server/internal/mysql/mock"
)

type RoomTestCase struct {
	name  string
	run   func(*testing.T, Room) error
	mock  func(room *mock_mysql.MockRoom)
	tx    func(sqlmock.Sqlmock, sqlmock.Sqlmock)
	check Check
}

type RoomTestCases []RoomTestCase

func (cs RoomTestCases) Run(t *testing.T) {
	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			mockDB, err := GetMockDB(t)
			if err != nil {
				t.Fatal(err)
			}
			defer mockDB.Close()

			c.tx(mockDB.rwm, mockDB.rom)

			mockRoom := mock_mysql.NewMockRoom(mockDB.ctrl)
			if c.mock != nil {
				c.mock(mockRoom)
			}

			repo := DefaultRoom{
				rw:   mockDB.rwgdb,
				ro:   mockDB.rogdb,
				Room: mockRoom,
			}

			err = c.run(t, &repo)
			c.check(t, err)

			err = mockDB.ExpectationsWereMet()
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestDefaultRoom_GetRoom(t *testing.T) {
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

	testCases := RoomTestCases{
		{
			name: "OK",
			run: func(t *testing.T, r Room) error {
				res, err := r.GetRoom(1)

				if diff := cmp.Diff(&value, res); len(diff) != 0 {
					t.Errorf("diff=%s", diff)
				}

				return err
			},
			mock: func(mock *mock_mysql.MockRoom) {
				mock.EXPECT().
					Find(gomock.Any(), 1).
					Return(&value, nil)
			},
			tx:    ReadOnly,
			check: Succeeded,
		},
		{
			name: "リソースが存在しない場合",
			run: func(t *testing.T, r Room) error {
				_, err := r.GetRoom(1)
				return err
			},
			mock: func(mock *mock_mysql.MockRoom) {
				mock.EXPECT().
					Find(gomock.Any(), 1).
					Return(nil, mysql.ErrNotFound)
			},
			tx:    ReadOnlyRollback,
			check: Failed(mysql.ErrNotFound),
		},
		{
			name: "Findがなんらかのエラーで失敗した場合",
			run: func(t *testing.T, r Room) error {
				_, err := r.GetRoom(1)
				return err
			},
			mock: func(mock *mock_mysql.MockRoom) {
				mock.EXPECT().
					Find(gomock.Any(), 1).
					Return(nil, errors.New(""))
			},
			tx:    ReadOnlyRollback,
			check: FailedAny,
		},
	}
	testCases.Run(t)
}
