package ast

import (
	"strconv"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/scroll"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

func MakeIdent(id token.Lexeme) Ident {
	return Ident{
		Base: Base{Snip: id.Snippet},
		Lex:  id,
	}
}

func MakeLiteral(lit token.Lexeme) Expr {
	switch lit.Token {
	case token.BOOL:
		return BoolLit{
			Base: Base{Snip: lit.Snippet},
			Val:  lit.Snippet.Text == "true",
		}

	case token.NUM:
		n, e := strconv.ParseFloat(lit.Snippet.Text, 64)
		if e != nil {
			panic("Illegal float format: '" + lit.Snippet.Text + "'")
		}
		return NumLit{
			Base: Base{Snip: lit.Snippet},
			Val:  n,
		}

	case token.STR:
		return StrLit{
			Base: Base{Snip: lit.Snippet},
			Val:  lit.Snippet.Text,
		}

	default:
		panic("Non-literal passed to makeLit")
	}
}

func MakeBinder(ids []Ident, op token.Lexeme, exprs []Expr) Binder {
	base := Base{
		Snip: scroll.Snippet{
			Start: ids[0].Snippet().Start,
			End:   exprs[0].Snippet().End,
		},
	}

	switch op.Token {
	case token.DEFINE:
		return Define{
			BinderBase: BinderBase{
				Base:  base,
				Op:    op,
				Left:  ids,
				Right: exprs,
			},
		}

	case token.ASSIGN:
		return Assign{
			BinderBase: BinderBase{
				Base:  base,
				Op:    op,
				Left:  ids,
				Right: exprs,
			},
		}

	default:
		panic("Non-binder passed to makeBinder")
	}
}
