package parser

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/ast"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/scroll"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

func err(itr LexIterator, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", itr.Line(), m)
	return errors.New(m)
}

func errNode(n ast.Node, m string, args ...interface{}) error {
	m = fmt.Sprintf(m, args...)
	m = fmt.Sprintf("Line %d: %s", n.Snippet().Start.Line, m)
	return errors.New(m)
}

func makeIdent(id token.Lexeme) ast.Ident {
	return ast.Ident{
		Snip: id.Snippet,
		Lex:  id,
	}
}

func makeLit(lit token.Lexeme) ast.Lit {
	return ast.Lit{
		Snip: lit.Snippet,
		Lex:  lit,
	}
}

func makeAssign(
	ids []ast.Ident,
	op token.Lexeme,
	exprs []ast.Expr,
) ast.Assign {
	return ast.Assign{
		Snip: scroll.Snippet{
			Start: ids[0].Snippet().Start,
			End:   exprs[0].Snippet().End,
		},
		Op:     op,
		Idents: ids,
		Exprs:  exprs,
	}
}
