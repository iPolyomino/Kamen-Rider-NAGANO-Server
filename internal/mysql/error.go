package mysql

import "errors"

var (
	// ErrNotFound はレコードが存在しないエラーを表します。
	ErrNotFound = errors.New("レコードが存在しません")
)
