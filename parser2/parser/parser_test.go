package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/number"

	"github.com/PaulioRandall/scarlet-go/token2/lexeme"
	"github.com/PaulioRandall/scarlet-go/token2/position"
	"github.com/PaulioRandall/scarlet-go/token2/token"
	"github.com/PaulioRandall/scarlet-go/token2/tokentest"

	"github.com/stretchr/testify/require"
)

func positionLexemes(lexs ...lexeme.Lexeme) []lexeme.Lexeme {
	tm := position.TextMarker{}
	for i := 0; i < len(lexs); i++ {
		v := lexs[i].Val
		lexs[i].Snippet = tm.Snippet(v)
		tm.Advance(v)
	}
	return lexs
}

func requireNodes(t *testing.T, exp, act []Node) {
	require.Equal(t, len(exp), len(act))
	for i, n := range act {
		require.Equal(t, exp[i], n)
	}
}

func TestParse_SingleAssign(t *testing.T) {

	// x := 1
	in := positionLexemes(
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok(":=", token.ASSIGN),
		lexeme.MakeTok("1", token.NUMBER),
	)

	exp := []Node{
		SingleAssign{
			Snippet: position.Snippet{
				UTF8Pos: in[0].Snippet.UTF8Pos,
				End:     in[2].Snippet.End,
			},
			Left:  Ident{Snippet: in[0].Snippet, Val: "x"},
			Infix: in[1].Snippet,
			Right: NumLit{Snippet: in[2].Snippet, Val: number.New("1")},
		},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_MultiAssign(t *testing.T) {

	// x, y, z := true, 1, "abc"
	in := positionLexemes(
		lexeme.MakeTok("x", token.IDENT), // 0
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("y", token.IDENT), // 2
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("z", token.IDENT), // 4
		lexeme.MakeTok(":=", token.ASSIGN),
		lexeme.MakeTok("true", token.TRUE), // 6
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok("1", token.NUMBER), // 8
		lexeme.MakeTok(",", token.DELIM),
		lexeme.MakeTok(`"text"`, token.STRING), // 10
	)

	exp := []Node{
		MultiAssign{
			Snippet: position.Snippet{
				UTF8Pos: in[0].Snippet.UTF8Pos,
				End:     in[10].Snippet.End,
			},
			Left: []Expr{
				Ident{Snippet: in[0].Snippet, Val: "x"},
				Ident{Snippet: in[2].Snippet, Val: "y"},
				Ident{Snippet: in[4].Snippet, Val: "z"},
			},
			Infix: in[5].Snippet,
			Right: []Expr{
				BoolLit{Snippet: in[6].Snippet, Val: true},
				NumLit{Snippet: in[8].Snippet, Val: number.New("1")},
				StrLit{Snippet: in[10].Snippet, Val: `"text"`},
			},
		},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_BinaryExpr_1(t *testing.T) {

	// 1 + 2
	in := positionLexemes(
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok("2", token.NUMBER),
	)

	exp := []Node{
		BinaryExpr{
			Snippet: position.Snippet{
				UTF8Pos: in[0].Snippet.UTF8Pos,
				End:     in[2].Snippet.End,
			},
			Left:  NumLit{Snippet: in[0].Snippet, Val: number.New("1")},
			Op:    in[1].Token,
			OpPos: in[1].Snippet,
			Right: NumLit{Snippet: in[2].Snippet, Val: number.New("2")},
		},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_BinaryExpr_2(t *testing.T) {

	// true && false
	in := positionLexemes(
		lexeme.MakeTok("true", token.TRUE),
		lexeme.MakeTok("&&", token.AND),
		lexeme.MakeTok("false", token.FALSE),
	)

	exp := []Node{
		BinaryExpr{
			Snippet: position.Snippet{
				UTF8Pos: in[0].Snippet.UTF8Pos,
				End:     in[2].Snippet.End,
			},
			Left:  BoolLit{Snippet: in[0].Snippet, Val: true},
			Op:    in[1].Token,
			OpPos: in[1].Snippet,
			Right: BoolLit{Snippet: in[2].Snippet, Val: false},
		},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_BinaryExpr_3(t *testing.T) {

	// 1 + 2 - 3
	in := positionLexemes(
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok("2", token.NUMBER),
		lexeme.MakeTok("-", token.SUB),
		lexeme.MakeTok("3", token.NUMBER),
	)

	add := BinaryExpr{
		Snippet: position.Snippet{
			UTF8Pos: in[0].Snippet.UTF8Pos,
			End:     in[2].Snippet.End,
		},
		Left:  NumLit{Snippet: in[0].Snippet, Val: number.New("1")},
		Op:    in[1].Token,
		OpPos: in[1].Snippet,
		Right: NumLit{Snippet: in[2].Snippet, Val: number.New("2")},
	}

	exp := []Node{
		BinaryExpr{
			Snippet: position.Snippet{
				UTF8Pos: in[0].Snippet.UTF8Pos,
				End:     in[4].Snippet.End,
			},
			Left:  add,
			Op:    in[3].Token,
			OpPos: in[3].Snippet,
			Right: NumLit{Snippet: in[4].Snippet, Val: number.New("3")},
		},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_BinaryExpr_4(t *testing.T) {

	// 1 + 2 * 3
	in := positionLexemes(
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok("2", token.NUMBER),
		lexeme.MakeTok("*", token.MUL),
		lexeme.MakeTok("3", token.NUMBER),
	)

	mul := BinaryExpr{
		Snippet: position.Snippet{
			UTF8Pos: in[2].Snippet.UTF8Pos,
			End:     in[4].Snippet.End,
		},
		Left:  NumLit{Snippet: in[2].Snippet, Val: number.New("2")},
		Op:    in[3].Token,
		OpPos: in[3].Snippet,
		Right: NumLit{Snippet: in[4].Snippet, Val: number.New("3")},
	}

	exp := []Node{
		BinaryExpr{
			Snippet: position.Snippet{
				UTF8Pos: in[0].Snippet.UTF8Pos,
				End:     in[4].Snippet.End,
			},
			Left:  NumLit{Snippet: in[0].Snippet, Val: number.New("1")},
			Op:    in[1].Token,
			OpPos: in[1].Snippet,
			Right: mul,
		},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_BinaryExpr_5(t *testing.T) {

	// 1 + 2 * 3 - 4 / 5 % 6
	// 1 + (2 * 3) - (4 / 5 % 6)
	in := positionLexemes(
		lexeme.MakeTok("1", token.NUMBER), // 0
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok("2", token.NUMBER), // 2
		lexeme.MakeTok("*", token.MUL),
		lexeme.MakeTok("3", token.NUMBER), // 4
		lexeme.MakeTok("-", token.SUB),
		lexeme.MakeTok("4", token.NUMBER), // 6
		lexeme.MakeTok("/", token.DIV),
		lexeme.MakeTok("5", token.NUMBER), // 8
		lexeme.MakeTok("%", token.REM),
		lexeme.MakeTok("6", token.NUMBER), // 10
	)

	// 2 * 3
	mul := BinaryExpr{
		Snippet: position.Snippet{
			UTF8Pos: in[2].Snippet.UTF8Pos,
			End:     in[4].Snippet.End,
		},
		Left:  NumLit{Snippet: in[2].Snippet, Val: number.New("2")},
		Op:    in[3].Token,
		OpPos: in[3].Snippet,
		Right: NumLit{Snippet: in[4].Snippet, Val: number.New("3")},
	}

	// 1 + (2 * 3)
	add := BinaryExpr{
		Snippet: position.Snippet{
			UTF8Pos: in[0].Snippet.UTF8Pos,
			End:     in[4].Snippet.End,
		},
		Left:  NumLit{Snippet: in[0].Snippet, Val: number.New("1")},
		Op:    in[1].Token,
		OpPos: in[1].Snippet,
		Right: mul,
	}

	// 4 / 5
	div := BinaryExpr{
		Snippet: position.Snippet{
			UTF8Pos: in[6].Snippet.UTF8Pos,
			End:     in[8].Snippet.End,
		},
		Left:  NumLit{Snippet: in[6].Snippet, Val: number.New("4")},
		Op:    in[7].Token,
		OpPos: in[7].Snippet,
		Right: NumLit{Snippet: in[8].Snippet, Val: number.New("5")},
	}

	// (4 / 5) % 6
	rem := BinaryExpr{
		Snippet: position.Snippet{
			UTF8Pos: in[6].Snippet.UTF8Pos,
			End:     in[10].Snippet.End,
		},
		Left:  div,
		Op:    in[9].Token,
		OpPos: in[9].Snippet,
		Right: NumLit{Snippet: in[10].Snippet, Val: number.New("6")},
	}

	exp := []Node{
		// (1 + 2 * 3) - (4 / 5 % 6)
		BinaryExpr{
			Snippet: position.Snippet{
				UTF8Pos: in[0].Snippet.UTF8Pos,
				End:     in[10].Snippet.End,
			},
			Left:  add,
			Op:    in[5].Token,
			OpPos: in[5].Snippet,
			Right: rem,
		},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_Assign_BinaryExpr_1(t *testing.T) {

	// x := 1 + 2
	in := positionLexemes(
		lexeme.MakeTok("x", token.IDENT),
		lexeme.MakeTok(":=", token.ASSIGN),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok("2", token.NUMBER),
	)

	right := BinaryExpr{
		Snippet: position.Snippet{
			UTF8Pos: in[2].Snippet.UTF8Pos,
			End:     in[4].Snippet.End,
		},
		Left:  NumLit{Snippet: in[2].Snippet, Val: number.New("1")},
		Op:    in[3].Token,
		OpPos: in[3].Snippet,
		Right: NumLit{Snippet: in[4].Snippet, Val: number.New("2")},
	}

	exp := []Node{
		SingleAssign{
			Snippet: position.Snippet{
				UTF8Pos: in[0].Snippet.UTF8Pos,
				End:     in[4].Snippet.End,
			},
			Left:  Ident{Snippet: in[0].Snippet, Val: "x"},
			Infix: in[1].Snippet,
			Right: right,
		},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_ParenExpr_1(t *testing.T) {

	// (1)
	in := positionLexemes(
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok(")", token.R_PAREN),
	)

	exp := []Node{
		NumLit{Snippet: in[1].Snippet, Val: number.New("1")},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_ParenExpr_2(t *testing.T) {

	// (1 + 2)
	in := positionLexemes(
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok("2", token.NUMBER),
		lexeme.MakeTok(")", token.R_PAREN),
	)

	exp := []Node{
		BinaryExpr{
			Snippet: position.Snippet{
				UTF8Pos: in[1].Snippet.UTF8Pos,
				End:     in[3].Snippet.End,
			},
			Left:  NumLit{Snippet: in[1].Snippet, Val: number.New("1")},
			Op:    in[2].Token,
			OpPos: in[2].Snippet,
			Right: NumLit{Snippet: in[3].Snippet, Val: number.New("2")},
		},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_ParenExpr_3(t *testing.T) {

	// ((1 + 2))
	in := positionLexemes(
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok("2", token.NUMBER),
		lexeme.MakeTok(")", token.R_PAREN),
		lexeme.MakeTok(")", token.R_PAREN),
	)

	exp := []Node{
		BinaryExpr{
			Snippet: position.Snippet{
				UTF8Pos: in[2].Snippet.UTF8Pos,
				End:     in[4].Snippet.End,
			},
			Left:  NumLit{Snippet: in[2].Snippet, Val: number.New("1")},
			Op:    in[3].Token,
			OpPos: in[3].Snippet,
			Right: NumLit{Snippet: in[4].Snippet, Val: number.New("2")},
		},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_ParenExpr_4(t *testing.T) {

	// ((1 + 2) * 3) - 4
	in := positionLexemes(
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("(", token.L_PAREN),
		lexeme.MakeTok("1", token.NUMBER),
		lexeme.MakeTok("+", token.ADD),
		lexeme.MakeTok("2", token.NUMBER),
		lexeme.MakeTok(")", token.R_PAREN),
		lexeme.MakeTok("*", token.MUL),
		lexeme.MakeTok("3", token.NUMBER),
		lexeme.MakeTok(")", token.R_PAREN),
		lexeme.MakeTok("-", token.SUB),
		lexeme.MakeTok("4", token.NUMBER),
	)

	// (1 + 2)
	add := BinaryExpr{
		Snippet: position.Snippet{
			UTF8Pos: in[2].Snippet.UTF8Pos,
			End:     in[4].Snippet.End,
		},
		Left:  NumLit{Snippet: in[2].Snippet, Val: number.New("1")},
		Op:    in[3].Token,
		OpPos: in[3].Snippet,
		Right: NumLit{Snippet: in[4].Snippet, Val: number.New("2")},
	}

	// ((1 + 2) * 3)
	mul := BinaryExpr{
		Snippet: position.Snippet{
			UTF8Pos: in[2].Snippet.UTF8Pos,
			End:     in[7].Snippet.End,
		},
		Left:  add,
		Op:    in[6].Token,
		OpPos: in[6].Snippet,
		Right: NumLit{Snippet: in[7].Snippet, Val: number.New("3")},
	}

	exp := []Node{
		// ((1 + 2) * 3) - 4
		BinaryExpr{
			Snippet: position.Snippet{
				UTF8Pos: in[2].Snippet.UTF8Pos,
				End:     in[10].Snippet.End,
			},
			Left:  mul,
			Op:    in[9].Token,
			OpPos: in[9].Snippet,
			Right: NumLit{Snippet: in[10].Snippet, Val: number.New("4")},
		},
	}

	tokenItr := tokentest.FeignSeries(in...)
	act, e := Parse(tokenItr)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}
