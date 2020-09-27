package scanner2

type reader struct {
	data []rune
	size int
	idx  int
	line int
	col  int
}

func (r *reader) more() bool {
	return r.idx >= r.size
}
