package parser

// IDs represents a set of IDs.
type IDs []string

// Params represents a set of parameters each need to be evaluated first.
type Params []Expr

// Block represents a code block where each expression must be evaluated in
// order.
type Block []Expr

// Spell represents a callable spell.
type Spell Params

// Assign represents an assignment.
type Assign struct {
	Dst IDs
	Src Params
}

// Func represents a callable function.
type Func struct {
	Params Params
	IDs    IDs
	Body   Block
}
