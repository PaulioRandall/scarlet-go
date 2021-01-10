package parser

import (
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
	//	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

type (
	// ParseTree is a recursion based tokeniser. It returns an AST and another
	// ParseTree function to obtain the following AST. On error or while
	// obtaining the last AST, ParseTree will be nil.
	ParseTree func() (ast.Tree, ParseTree, error)
)

// ParseAll parses all lexemes into a slice of ASTs.
func ParseAll(itr LexIterator) []ast.Tree {
	// TODO
	return nil
}

// New returns a new ParseTree function.
func New(itr LexIterator) ParseTree {
	// TODO
	return nil
}
