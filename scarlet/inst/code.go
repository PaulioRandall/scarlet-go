package inst

// Code represents a high level instruction for an interpreter.
type Code int

const (
	UNDEFINED Code = iota

	// Push a specified value from the static data pool on to the top of the
	// value stack.
	STACK_PUSH

	// Pop a value off the top of the value stack and discard it.
	STACK_POP

	// Fetches a specified value from the stack and pushes it on to the top of the
	// value stack.
	FETCH_PUSH

	// Pop a value off the top of the value stack and bind it to an identifier
	// within the current scope.
	SCOPE_BIND

	// Get the value bound to a specific identifier.
	//SCOPE_GET = "SCOPE_GET"

	// Pop two values off the top of the value stack, perform the specified
	// binary operation, then push the result onto the top of the value stack.

	BIN_OP_ADD
	BIN_OP_SUB
	BIN_OP_MUL
	BIN_OP_DIV
	BIN_OP_REM

	BIN_OP_AND
	BIN_OP_OR

	BIN_OP_LESS
	BIN_OP_MORE
	BIN_OP_LEQU
	BIN_OP_MEQU
	BIN_OP_EQU
	BIN_OP_NEQU
)

// String returns a human readable string representation of the Code.
func (c Code) String() string {
	switch c {
	case UNDEFINED:
		return "Undefined"
	case STACK_PUSH:
		return "STACK_PUSH"
	case STACK_POP:
		return "STACK_POP"
	case SCOPE_BIND:
		return "SCOPE_BIND"

	case BIN_OP_ADD:
		return "BIN_OP_ADD"
	case BIN_OP_SUB:
		return "BIN_OP_SUB"
	case BIN_OP_MUL:
		return "BIN_OP_MUL"
	case BIN_OP_DIV:
		return "BIN_OP_DIV"
	case BIN_OP_REM:
		return "BIN_OP_REM"

	case BIN_OP_AND:
		return "BIN_OP_AND"
	case BIN_OP_OR:
		return "BIN_OP_OR"

	case BIN_OP_LESS:
		return "BIN_OP_LESS"
	case BIN_OP_MORE:
		return "BIN_OP_MORE"
	case BIN_OP_LEQU:
		return "BIN_OP_LEQU"
	case BIN_OP_MEQU:
		return "BIN_OP_MEQU"
	case BIN_OP_EQU:
		return "BIN_OP_EQU"
	case BIN_OP_NEQU:
		return "BIN_OP_NEQU"

	default:
		return ""
	}
}
