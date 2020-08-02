package inst

type Code int

func (c Code) String() string {
	return codes[c]
}

const (
	CO_UNDEFINED Code = iota

	// Set the program counter to the instruction number at the top of the value
	// stack.
	//IN_JUMP

	// Push the instruction number at the top of the value stack onto the call
	// stack.
	//IN_REF_PUSH

	// Pop an instruction off the call stack and onto the value stack.
	//IN_REF_POP

	// Push a value onto the value stack.
	//
	// Data: value
	// Stack values produced: 1
	//    1st: value
	CO_VAL_PUSH

	// Pop a value off the value stack and put it into the current context.
	//
	// Expected value stack:
	// top: identifier
	//IN_CTX_SET

	// Queries a value from the current context and push it onto the value stack.
	//
	// Data: identifier
	// Stack values produced: 1
	//    1st: query response value
	CO_CTX_GET

	// Call a spell with the contents of the value stack.
	//
	// Data: spell name
	// Stack values consumed: 1+
	//    1st: number of spell arguments
	//    ...: spell input
	// Stack values produced: 0+
	//    ...: spell output
	CO_SPELL
)

// Example: @Println("Scarlet")
// 1: IN_VAL_PUSH   "Scarlet"
// 2: IN_SPELL      @Println

// Example: @Exit(0)
// 1: IN_VAL_PUSH   0
// 2: IN_SPELL      @Exit

// Example: @Set("x", "Scarlet")
// 1: IN_VAL_PUSH   "x"
// 2: IN_VAL_PUSH   "Scarlet"
// 2: IN_SPELL      @Set

// Example: @Set("x", y)
// 1: IN_VAL_PUSH   "x"
// 2: IN_CTX_GET    y
// 2: IN_SPELL      @Set

var codes = map[Code]string{
	CO_VAL_PUSH: `CO_VAL_PUSH`,
	CO_CTX_GET:  `CO_CTX_GET`,
	CO_SPELL:    `CO_SPELL`,
}