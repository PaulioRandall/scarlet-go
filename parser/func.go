package parser

// IDs represents a set of IDs.
type IDs []string

// Params represents a set of parameters each need to be evaluated first.
type Params []Expr

// Block represents a code block where each expression must be evaluated in
// order.
type Block []Expr

// Func represents a callable function.
type Func struct {
	Params Params
	IDs    IDs
	Body   Block
}
