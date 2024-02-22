package graph

import "errors"

var (
	ErrSomethingWrong = errors.New("something went wrong")
	ErrNotFound       = errors.New("not found")
	ErrCannotDelete   = errors.New("cannot delete this data")
)

var (
	BaseOffset int32 = 0
	BaseLimit  int32 = 0
)
