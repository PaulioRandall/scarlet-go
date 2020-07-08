package runtime

type resultType int

const (
	RT_UNDEFINED resultType = iota
	RT_BOOL
	RT_NUMBER
	RT_STRING
)

var resultTypes = map[resultType]string{
	RT_BOOL:   "bool",
	RT_NUMBER: "number",
	RT_STRING: "string",
}

func (rt resultType) String() string {
	return resultTypes[rt]
}

type result struct {
}
