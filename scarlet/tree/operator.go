package tree

type Operator int

const (
	OP_UNDEFINED Operator = iota
	OP_ADD                // +
	OP_SUB                // -
	OP_MUL                // *
	OP_DIV                // /
	OP_REM                // %
	OP_AND                // &&
	OP_OR                 // ||
	OP_LT                 // <
	OP_MT                 // >
	OP_LTE                // <=
	OP_MTE                // >=
	OP_EQU                // ==
	OP_NEQ                // !=
	OP_EXIST              // ?
)
