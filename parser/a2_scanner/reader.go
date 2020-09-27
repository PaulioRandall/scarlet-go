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

func (r *reader) at(offset int) rune {
	return r.data[offset+r.idx]
}

func (r *reader) starts(s string) bool {
	return r.contains(0, s)
}

func (r *reader) contains(start int, s string) bool {

	i := r.idx + start
	for _, ru := range s {
		if r.data[i] != ru {
			return false
		}
		i++
	}

	return true
}
