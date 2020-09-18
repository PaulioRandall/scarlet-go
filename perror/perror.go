package perror

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
