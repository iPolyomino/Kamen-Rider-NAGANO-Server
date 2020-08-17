package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

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
		t.Errorf("expected no error but actual: %s", err)
	}
}

func ReadOnly(_ sqlmock.Sqlmock, ro sqlmock.Sqlmock) {
	ro.ExpectBegin()
	ro.ExpectCommit()
}

func ReadOnlyRollback(_ sqlmock.Sqlmock, ro sqlmock.Sqlmock) {
	ro.ExpectBegin()
	ro.ExpectRollback()
}
