package perror

import (
	"fmt"

	"github.com/pkg/errors"
)

func New(msg string, args ...interface{}) error {
	e := fmt.Errorf(msg, args...)
	return errors.WithStack(e)
}

func Panic(msg string, args ...interface{}) error {
	e := fmt.Errorf(msg, args...)
	panic(errors.WithStack(e))
}
