package store

import "errors"

var ErrResultNotFound = errors.New("result not found")

type Store interface {
	Ping() error
}
