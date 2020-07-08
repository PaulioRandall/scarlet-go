package inst

type Code int

const (
	IN_UNDEFINED Code = iota

	// Set the program counter to the instruction number at the top of the value
	// stack.
	//IN_JUMP

	// Push the instruction number at the top of the value stack onto the call
	// stack.
	//IN_REF_PUSH

	// Pop an instruction off the call stack and onto the value stack.
	//IN_REF_POP

	// Push the item supplied with the instruction onto the value stack.
	//
	// Arg: item
	IN_VAL_PUSH

	// Pop a value off the value stack and put it into the current context.
	//
	// Expected value stack:
	// top: identifier
	//IN_CTX_SET

	// Gets a variable from the current context and push it onto the value stack.
	//
	// Arg: identifier
	IN_CTX_GET

	// Call a spell with the contents of the value stack.
	//
	// Arg: spell name
	//
	// Expected value stack:
	// 001: number of arguments
	// ...: args
	IN_SPELL
)

// Example: @Println("Scarlet")
// 1: IN_VAL_PUSH		"Scarlet"
// 2: IN_SPELL   		@Println

// Example: @Exit(0)
// 1: IN_VAL_PUSH		0
// 2: IN_SPELL   		@Exit

// Example: @Set("x", "Scarlet")
// 1: IN_VAL_PUSH		"x"
// 2: IN_VAL_PUSH		"Scarlet"
// 2: IN_SPELL   		@Set

// Example: @Set("x", y)
// 1: IN_VAL_PUSH		"x"
// 2: IN_CTX_GET 		y
// 2: IN_SPELL   		@Set
