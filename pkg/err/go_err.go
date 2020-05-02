package err

// goErr wraps a Go error as an Err.
type goErr struct {
	error
}

func (ge goErr) Error() string {
	return ge.error.Error()
}

func (ge goErr) Cause() error {
	return nil
}

func (ge goErr) LineIndex() int {
	return 0
}

func (ge goErr) ColIndex() int {
	return 0
}

func (ge goErr) Length() int {
	return 0
}
