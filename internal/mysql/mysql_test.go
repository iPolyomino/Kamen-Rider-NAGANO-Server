package mysql

import (
	"errors"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/DATA-DOG/go-sqlmock"
)

type TestCases []TestCase

type TestCase struct {
	name  string
	run   func(*testing.T, *gorm.DB) error
	mock  func(sqlmock.Sqlmock)
	check Check
}

func (cs TestCases) Run(t *testing.T) {
	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			if c.mock != nil {
				c.mock(mock)
			}

			gdb, err := gorm.Open("mysql", db)
			if err != nil {
				t.Fatal(err)
			}
			defer gdb.Close()

			err = gdb.Transaction(func(tx *gorm.DB) error {
				return c.run(t, tx)
			})
			c.check(t, err)

			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Error(err)
			}
		})
	}
}

type Check func(*testing.T, error)

func Failed(err error) Check {
	return func(t *testing.T, e error) {
		if !errors.Is(e, err) {
			t.Errorf("expected %v but actual %v", err, e)
		}
	}
}

func FailedAny(t *testing.T, err error) {
	if err == nil {
		t.Error("no error")
	}
}

func Succeeded(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expected no error: %s", err)
	}
}
