package parser

type err struct {
	msg string
}

func (e err) Error() string {
	return e.msg
}

func newErr(msg string) error {
	return err{msg: msg}
}
