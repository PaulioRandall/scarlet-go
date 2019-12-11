package parser2

type EvalErr interface {
	Error() string

	Who() int64
}

type stdEvalErr struct {
	what     string
	tokenRef int64
}

func (e stdEvalErr) Error() string {
	return e.what
}

func (e stdEvalErr) Who() int64 {
	return e.tokenRef
}
