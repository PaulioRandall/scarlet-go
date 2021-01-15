package ast

import (
	"strconv"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/scroll"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

func MakeIdent(id token.Lexeme) Ident {
	return Ident{
		BaseExpr: BaseExpr{
			ValType: T_INFER,
			Base:    Base{Snip: id.Snippet},
		},
		Lex: id,
	}
}

func MakeLiteral(lit token.Lexeme) Expr {
	switch val := lit.Snippet.Text; lit.Token {
	case token.BOOL:
		return BoolLit{
			BaseExpr: BaseExpr{
				ValType: T_BOOL,
				Base:    Base{Snip: lit.Snippet},
			},
			Val: val == "true",
		}

	case token.NUM:
		n, e := strconv.ParseFloat(val, 64)
		if e != nil {
			panic("Illegal float format: '" + val + "'")
		}
		return NumLit{
			BaseExpr: BaseExpr{
				ValType: T_NUM,
				Base:    Base{Snip: lit.Snippet},
			},
			Val: n,
		}

	case token.STR:
		return StrLit{
			BaseExpr: BaseExpr{
				ValType: T_STR,
				Base:    Base{Snip: lit.Snippet},
			},
			Val: val[1 : len(val)-1], // Remove double quotes
		}

	default:
		panic("Non-literal passed to makeLit")
	}
}

func MakeBinding(ids []Ident, op token.Lexeme, exprs []Expr) Binding {
	base := Base{
		Snip: scroll.Snippet{
			Start: ids[0].Snippet().Start,
			End:   exprs[0].Snippet().End,
		},
	}

	switch op.Token {
	case token.DEFINE:
		return Define{
			BaseBinding: BaseBinding{
				Base:  base,
				Op:    op,
				Left:  ids,
				Right: exprs,
			},
		}

	case token.ASSIGN:
		return Assign{
			BaseBinding: BaseBinding{
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
