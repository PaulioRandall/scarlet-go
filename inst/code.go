package inst

type Code int

func (c Code) String() string {
	return codes[c]
}

const (
	CO_UNDEFINED Code = iota

	// Decrement the instruction counter by the amount provided as data.
	//
	// Data: integer
	CO_JMP_BACK

	// Pops a value off the stack, which should be a bool value, if it's false
	// then increment the instruction counter by the amount provided as data.
	//
	// Data: integer
	// Stack values consumed: 1
	// 		1st: bool
	CO_JMP_FALSE

	// Push a delimiter onto the value stack. Delimiters are used to indicate
	// when to stop popping values of the value stack for instruction that
	// have variable length arguments, e.g. spell arguments and assignments.
	//
	// Stack values produced: 1
	//    1st: delimiter
	CO_DLM_PUSH

	// Push a value onto the value stack.
	//
	// Data: value
	// Stack values produced: 1
	//    1st: any value type
	CO_VAL_PVAL // PUSH_VAL

	// Push a nil value onto the value stack.
	//
	// Stack values produced: 1
	//    1st: any value type
	CO_VAL_PNIL // PUSH_NIL

	// Pop a value off the value stack and bind it to an identifier within the
	// the current context. If the value is a nil then an unbinding should
	// occur instead.
	//
	// Data: identifier
	// Stack values consumed: 1
	// 		1st: any type
	CO_VAL_BIND

	// Queries a value from the current context and push it onto the value stack.
	//
	// Data: identifier
	// Stack values produced: 1
	//    1st: query response value
	CO_VAL_GET

	// Pop a value off the value stack and throw it away.
	//
	// Stack values consumed: 1
	// 		1st: any type
	CO_VAL_POP

	// Pushes a new sub-context into the current context's sub-context call stack.
	CO_SUB_CTX_PUSH

	// Pops a top sub-context from the current context's sub-context call stack.
	CO_SUB_CTX_POP

	// Performs an arithmetic operation on two operands and pushes the result
	// onto the value stack.
	//
	// Stack values consumed: 2
	//    1st: right operand
	//    2nd: left operand
	// Stack values produced: 1
	//    1st: operation result
	CO_MUL
	CO_DIV
	CO_REM
	CO_ADD
	CO_SUB
	CO_AND
	CO_OR
	CO_LESS
	CO_MORE
	CO_LESS_EQU
	CO_MORE_EQU
	CO_EQU
	CO_NOT_EQU

	// Call a spell with the contents of the value stack.
	//
	// Data: spell name
	// Stack values consumed: 1+
	//    1st: number of spell arguments
	//    ...: spell input
	// Stack values produced: 0+
	//    ...: spell output
	CO_SPL_CALL
)

var codes = map[Code]string{
	CO_JMP_BACK:     `CO_JMP_BACK`,
	CO_JMP_FALSE:    `CO_JMP_FALSE`,
	CO_DLM_PUSH:     `CO_DLM_PUSH`,
	CO_VAL_PVAL:     `CO_VAL_PVAL`,
	CO_VAL_PNIL:     `CO_VAL_PNIL`,
	CO_VAL_BIND:     `CO_VAL_BIND`,
	CO_VAL_GET:      `CO_VAL_GET`,
	CO_VAL_POP:      `CO_VAL_POP`,
	CO_SUB_CTX_PUSH: `CO_SUB_CTX_PUSH`,
	CO_SUB_CTX_POP:  `CO_SUB_CTX_POP`,
	CO_SPL_CALL:     `CO_SPL_CALL`,
	CO_MUL:          `CO_MUL`,
	CO_DIV:          `CO_DIV`,
	CO_REM:          `CO_REM`,
	CO_ADD:          `CO_ADD`,
	CO_SUB:          `CO_SUB`,
	CO_AND:          `CO_AND`,
	CO_OR:           `CO_OR`,
	CO_LESS:         `CO_LESS`,
	CO_MORE:         `CO_MORE`,
	CO_LESS_EQU:     `CO_LESS_EQU`,
	CO_MORE_EQU:     `CO_MORE_EQU`,
	CO_EQU:          `CO_EQU`,
	CO_NOT_EQU:      `CO_NOT_EQU`,
}
