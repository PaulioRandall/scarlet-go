package stream

type Stream struct {
	tks  []interface{}
	size int
	idx  int
}

func New(items []interface{}) *Stream {
	return &Stream{
		tks:  items,
		size: len(items),
	}
}

func (stm *Stream) Next() interface{} {

	if stm.idx >= stm.size {
		return nil
	}

	tk := stm.tks[stm.idx]
	stm.idx++
	return tk
}
