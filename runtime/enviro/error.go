package enviro

import (
	"fmt"
)

type enviroErr struct {
	error
	msg string
}

func newErr(msg string, args ...interface{}) error {
	return enviroErr{
		msg: fmt.Sprintf(msg, args...),
	}
}

func (e enviroErr) Error() string {
	return e.msg
}
