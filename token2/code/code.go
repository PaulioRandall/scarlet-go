package code

// Code represents a high level instruction for an interpreter.
type Code int

const (
	UNDEFINED Code = iota

	// Push a specified value from the static data pool on to the top of the
	// value stack.
	STACK_PUSH

	// Pop a value off the top of the value stack and discard it.
	STACK_POP

	// Pop a value off the top of the value stack and bind it to an identifier
	// within the current scope.
	SCOPE_BIND

	// Get the value bound to a specific identifier.
	//SCOPE_GET = "SCOPE_GET"

	// Pop two values off the top of the value stack, perform the specified
	// binary operation, then push the result onto the top of the value stack.

	OP_ADD
	OP_SUB
	OP_MUL
	OP_DIV
	OP_REM

	OP_AND
	OP_OR

	OP_LESS
	OP_MORE
	OP_LEQU
	OP_MEQU
	OP_EQU
	OP_NEQU
)

// String returns a human readable string representation of the Code.
func (c Code) String() string {
	switch c {
	case UNDEFINED:
		return "Undefined"
	case STACK_PUSH:
		return "Push_on_value_stack"
	case STACK_POP:
		return "Pop_off_value_stack"
	case SCOPE_BIND:
		return "Bind_to_identifier"

	case OP_ADD:
		return "Add"
	case OP_SUB:
		return "Subtract"
	case OP_MUL:
		return "Multiple"
	case OP_DIV:
		return "Divide"
	case OP_REM:
		return "Remainder"

	case OP_AND:
		return "Logical_and"
	case OP_OR:
		return "Logical_or"

	case OP_LESS:
		return "Less_than"
	case OP_MORE:
		return "More_than"
	case OP_LEQU:
		return "Less_than_or_equal"
	case OP_MEQU:
		return "More_than_or_equal"
	case OP_EQU:
		return "Equal"
	case OP_NEQU:
		return "Not_equal"

	default:
		return ""
	}
}
