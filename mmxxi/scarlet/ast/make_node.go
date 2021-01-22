package ast

import (
	"strconv"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/scroll"
	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/token"
)

func MakeVar(id token.Lexeme, t ValType) Var {
	return Var{
		Base:    Base{Snip: id.Snippet},
		ValType: t,
		Lex:     id,
	}
}

func MakeIdent(id token.Lexeme, t ValType) Ident {
	return Ident{
		BaseExpr: BaseExpr{
			ValType: t,
			Base:    Base{Snip: id.Snippet},
		},
		Val: id.Snippet.Text,
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

func MakeBinding(vars []Var, op token.Lexeme, exprs []Expr) Binding {
	base := Base{
		Snip: scroll.Snippet{
			Start: vars[0].Snippet().Start,
			End:   exprs[0].Snippet().End,
		},
	}

	bb := BaseBinding{
		Base:  base,
		Op:    op,
		Left:  vars,
		Right: exprs,
	}

	switch op.Token {
	case token.DEFINE:
		return Define{BaseBinding: bb}

	case token.ASSIGN:
		return Assign{BaseBinding: bb}

	default:
		panic("Non-binder passed to makeBinder")
	}
}
