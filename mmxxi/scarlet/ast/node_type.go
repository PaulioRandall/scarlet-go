package ast

import (
//"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/scroll"
)

// NodeType represents programming constructs comprised of token patterns.
type NodeType int

const (
	UNDEFINED NodeType = iota
	IDENT
	LITERAL
	DEFINE
	ASSIGN
)
