package result

type ResultType int

const (
	RT_UNDEFINED ResultType = iota
	RT_VOID
)

var types = map[ResultType]string{
	RT_VOID: "void",
}

func (rt ResultType) String() string {
	return types[rt]
}
