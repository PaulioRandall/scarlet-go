package perror

import (
	"fmt"
)

func New(msg string, args ...interface{}) error {
	e := fmt.Errorf(msg, args...)
	return e
}

func Panic(msg string, args ...interface{}) error {
	e := fmt.Errorf(msg, args...)
	panic(e)
}

type Perror interface {
	error
}

type Pos interface {
	Pos() (lineIdx, colIdx int)
}

type Len interface {
	Len() int
}

type Eof interface {
	Eof() bool
}
