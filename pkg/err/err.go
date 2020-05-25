package err

type Err interface {
	error
	Line() int // index
	Col() int  // index
	Len() int
}

type sErr struct {
	msg  string
	line int
	col  int
	len  int
}

func (e sErr) Error() string {
	return e.msg
}

func (e sErr) Line() int {
	return e.line
}

func (e sErr) Col() int {
	return e.col
}

func (e sErr) Len() int {
	return e.len
}

type Option func(*sErr)

type Lexeme interface {
	Value() string
	Line() int
	Col() int
}

func Panic(msg string, ops ...Option) {
	er := New(msg, ops...)
	panic(er)
}

func Epanic(e error, ops ...Option) {
	er := Wrap(e, ops...)
	panic(er)
}

func New(msg string, ops ...Option) Err {

	s := sErr{
		msg:  msg,
		line: -1,
		col:  -1,
	}

	applyOptions(&s, ops...)
	return s
}

func Wrap(e error, ops ...Option) Err {

	s := sErr{
		msg:  e.Error(),
		line: -1,
		col:  -1,
	}

	applyOptions(&s, ops...)
	return s
}

func applyOptions(e *sErr, ops ...Option) {
	for _, op := range ops {
		op(e)
	}
}

func Pos(line, col int) Option {
	return func(s *sErr) {
		s.line = line
		s.col = col
	}
}

func Len(len int) Option {
	return func(s *sErr) {
		s.len = len
	}
}

func At(lex Lexeme) Option {
	l, c, ln := lex.Line(), lex.Col(), len(lex.Value())

	return func(s *sErr) {
		s.line = l
		s.col = c
		s.len = ln
	}
}

func After(lex Lexeme) Option {
	l := lex.Line()
	c := lex.Col() + len(lex.Value())

	return func(s *sErr) {
		s.line = l
		s.col = c
		s.len = 0
	}
}

func Try(f func()) (err Err) {

	func() {
		defer func() {

			switch v := recover().(type) {
			case nil:
				err = nil
			case Err:
				err = v
			case string:
				err = New(v)
			case error:
				err = Wrap(v)
			default:
				s := `¯\_(ツ)_/¯ Something panicked, but I don't understand the error`
				err = New(s)
			}

		}()

		f()
	}()

	return
}
