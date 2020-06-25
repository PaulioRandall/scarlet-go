package runtime

type ResultType int

const (
	RT_UNDEFINED ResultType = iota
	RT_VOID
	RT_BOOL
	RT_NUMBER
	RT_STRING
	RT_LIST
	RT_MAP
	RT_FUNC_DEF
	RT_EXPR_FUNC_DEF
	RT_TUPLE
)

var types = map[ResultType]string{
	RT_VOID:          "void",
	RT_BOOL:          "bool",
	RT_NUMBER:        "number",
	RT_STRING:        "string",
	RT_LIST:          "list",
	RT_MAP:           "map",
	RT_FUNC_DEF:      "function",
	RT_EXPR_FUNC_DEF: "expression-function",
	RT_TUPLE:         "tuple",
}

func (rt ResultType) String() string {
	return types[rt]
}
