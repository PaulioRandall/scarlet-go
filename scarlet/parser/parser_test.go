package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"
	"github.com/PaulioRandall/scarlet-go/scarlet/value/number"

	"github.com/stretchr/testify/require"
)

func positionLexemes(lexs ...token.Lexeme) []token.Lexeme {
	tm := token.TextMarker{}
	for i := 0; i < len(lexs); i++ {
		v := lexs[i].Val
		lexs[i].Snippet = tm.Snippet(v)
		tm.Advance(v)
	}
	return lexs
}

func superSnip(start, end token.Lexeme) token.Snippet {
	return token.Snippet{
		UTF8Pos: start.Snippet.UTF8Pos,
		End:     end.Snippet.End,
	}
}

func requireNodes(t *testing.T, exp, act []tree.Node) {
	require.Equal(t, len(exp), len(act))
	for i, n := range act {
		require.Equal(t, exp[i], n)
	}
}

func TestParse_SingleAssign_1(t *testing.T) {

	// x := 1
	in := positionLexemes(
		token.MakeTok("x", token.IDENT),
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("1", token.NUMBER),
	)

	exp := []tree.Node{
		tree.SingleAssign{
			Snippet: superSnip(in[0], in[2]),
			Left:    tree.Ident{Snippet: in[0].Snippet, Val: "x"},
			Infix:   in[1].Snippet,
			Right:   tree.NumLit{Snippet: in[2].Snippet, Val: number.New("1")},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_SingleAssign_2(t *testing.T) {

	// x := 1
	in := positionLexemes(
		token.MakeTok("x", token.IDENT),
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("y", token.IDENT),
	)

	exp := []tree.Node{
		tree.SingleAssign{
			Snippet: superSnip(in[0], in[2]),
			Left:    tree.Ident{Snippet: in[0].Snippet, Val: "x"},
			Infix:   in[1].Snippet,
			Right:   tree.Ident{Snippet: in[2].Snippet, Val: "y"},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_MultiAssign(t *testing.T) {

	// x, y, z := true, 1, "abc"
	in := positionLexemes(
		token.MakeTok("x", token.IDENT), // 0
		token.MakeTok(",", token.DELIM),
		token.MakeTok("y", token.IDENT), // 2
		token.MakeTok(",", token.DELIM),
		token.MakeTok("z", token.IDENT), // 4
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("true", token.TRUE), // 6
		token.MakeTok(",", token.DELIM),
		token.MakeTok("1", token.NUMBER), // 8
		token.MakeTok(",", token.DELIM),
		token.MakeTok(`"text"`, token.STRING), // 10
	)

	exp := []tree.Node{
		tree.MultiAssign{
			Snippet: superSnip(in[0], in[10]),
			Left: []tree.Assignee{
				tree.Ident{Snippet: in[0].Snippet, Val: "x"},
				tree.Ident{Snippet: in[2].Snippet, Val: "y"},
				tree.Ident{Snippet: in[4].Snippet, Val: "z"},
			},
			Infix: in[5].Snippet,
			Right: []tree.Expr{
				tree.BoolLit{Snippet: in[6].Snippet, Val: true},
				tree.NumLit{Snippet: in[8].Snippet, Val: number.New("1")},
				tree.StrLit{Snippet: in[10].Snippet, Val: `"text"`},
			},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_IdentExpr_1(t *testing.T) {

	// x
	in := positionLexemes(
		token.MakeTok("x", token.IDENT),
	)

	exp := []tree.Node{
		tree.Ident{Snippet: in[0].Snippet, Val: "x"},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_BinaryExpr_1(t *testing.T) {

	// 1 + 2
	in := positionLexemes(
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("+", token.ADD),
		token.MakeTok("2", token.NUMBER),
	)

	exp := []tree.Node{
		tree.BinaryExpr{
			Snippet: superSnip(in[0], in[2]),
			Left:    tree.NumLit{Snippet: in[0].Snippet, Val: number.New("1")},
			Op:      in[1].Token,
			OpPos:   in[1].Snippet,
			Right:   tree.NumLit{Snippet: in[2].Snippet, Val: number.New("2")},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_BinaryExpr_2(t *testing.T) {

	// true && false
	in := positionLexemes(
		token.MakeTok("true", token.TRUE),
		token.MakeTok("&&", token.AND),
		token.MakeTok("false", token.FALSE),
	)

	exp := []tree.Node{
		tree.BinaryExpr{
			Snippet: superSnip(in[0], in[2]),
			Left:    tree.BoolLit{Snippet: in[0].Snippet, Val: true},
			Op:      in[1].Token,
			OpPos:   in[1].Snippet,
			Right:   tree.BoolLit{Snippet: in[2].Snippet, Val: false},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_BinaryExpr_3(t *testing.T) {

	// 1 + 2 - 3
	in := positionLexemes(
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("+", token.ADD),
		token.MakeTok("2", token.NUMBER),
		token.MakeTok("-", token.SUB),
		token.MakeTok("3", token.NUMBER),
	)

	add := tree.BinaryExpr{
		Snippet: superSnip(in[0], in[2]),
		Left:    tree.NumLit{Snippet: in[0].Snippet, Val: number.New("1")},
		Op:      in[1].Token,
		OpPos:   in[1].Snippet,
		Right:   tree.NumLit{Snippet: in[2].Snippet, Val: number.New("2")},
	}

	exp := []tree.Node{
		tree.BinaryExpr{
			Snippet: superSnip(in[0], in[4]),
			Left:    add,
			Op:      in[3].Token,
			OpPos:   in[3].Snippet,
			Right:   tree.NumLit{Snippet: in[4].Snippet, Val: number.New("3")},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_BinaryExpr_4(t *testing.T) {

	// 1 + 2 * 3
	in := positionLexemes(
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("+", token.ADD),
		token.MakeTok("2", token.NUMBER),
		token.MakeTok("*", token.MUL),
		token.MakeTok("3", token.NUMBER),
	)

	mul := tree.BinaryExpr{
		Snippet: superSnip(in[2], in[4]),
		Left:    tree.NumLit{Snippet: in[2].Snippet, Val: number.New("2")},
		Op:      in[3].Token,
		OpPos:   in[3].Snippet,
		Right:   tree.NumLit{Snippet: in[4].Snippet, Val: number.New("3")},
	}

	exp := []tree.Node{
		tree.BinaryExpr{
			Snippet: superSnip(in[0], in[4]),
			Left:    tree.NumLit{Snippet: in[0].Snippet, Val: number.New("1")},
			Op:      in[1].Token,
			OpPos:   in[1].Snippet,
			Right:   mul,
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_BinaryExpr_5(t *testing.T) {

	// 1 + 2 * 3 - 4 / 5 % 6
	// 1 + (2 * 3) - (4 / 5 % 6)
	in := positionLexemes(
		token.MakeTok("1", token.NUMBER), // 0
		token.MakeTok("+", token.ADD),
		token.MakeTok("2", token.NUMBER), // 2
		token.MakeTok("*", token.MUL),
		token.MakeTok("3", token.NUMBER), // 4
		token.MakeTok("-", token.SUB),
		token.MakeTok("4", token.NUMBER), // 6
		token.MakeTok("/", token.DIV),
		token.MakeTok("5", token.NUMBER), // 8
		token.MakeTok("%", token.REM),
		token.MakeTok("6", token.NUMBER), // 10
	)

	// 2 * 3
	mul := tree.BinaryExpr{
		Snippet: superSnip(in[2], in[4]),
		Left:    tree.NumLit{Snippet: in[2].Snippet, Val: number.New("2")},
		Op:      in[3].Token,
		OpPos:   in[3].Snippet,
		Right:   tree.NumLit{Snippet: in[4].Snippet, Val: number.New("3")},
	}

	// 1 + (2 * 3)
	add := tree.BinaryExpr{
		Snippet: superSnip(in[0], in[4]),
		Left:    tree.NumLit{Snippet: in[0].Snippet, Val: number.New("1")},
		Op:      in[1].Token,
		OpPos:   in[1].Snippet,
		Right:   mul,
	}

	// 4 / 5
	div := tree.BinaryExpr{
		Snippet: superSnip(in[6], in[8]),
		Left:    tree.NumLit{Snippet: in[6].Snippet, Val: number.New("4")},
		Op:      in[7].Token,
		OpPos:   in[7].Snippet,
		Right:   tree.NumLit{Snippet: in[8].Snippet, Val: number.New("5")},
	}

	// (4 / 5) % 6
	rem := tree.BinaryExpr{
		Snippet: superSnip(in[6], in[10]),
		Left:    div,
		Op:      in[9].Token,
		OpPos:   in[9].Snippet,
		Right:   tree.NumLit{Snippet: in[10].Snippet, Val: number.New("6")},
	}

	exp := []tree.Node{
		// (1 + 2 * 3) - (4 / 5 % 6)
		tree.BinaryExpr{
			Snippet: superSnip(in[0], in[10]),
			Left:    add,
			Op:      in[5].Token,
			OpPos:   in[5].Snippet,
			Right:   rem,
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_Assign_BinaryExpr_1(t *testing.T) {

	// x := 1 + 2
	in := positionLexemes(
		token.MakeTok("x", token.IDENT),
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("+", token.ADD),
		token.MakeTok("2", token.NUMBER),
	)

	right := tree.BinaryExpr{
		Snippet: superSnip(in[2], in[4]),
		Left:    tree.NumLit{Snippet: in[2].Snippet, Val: number.New("1")},
		Op:      in[3].Token,
		OpPos:   in[3].Snippet,
		Right:   tree.NumLit{Snippet: in[4].Snippet, Val: number.New("2")},
	}

	exp := []tree.Node{
		tree.SingleAssign{
			Snippet: superSnip(in[0], in[4]),
			Left:    tree.Ident{Snippet: in[0].Snippet, Val: "x"},
			Infix:   in[1].Snippet,
			Right:   right,
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_ParenExpr_1(t *testing.T) {

	// (1)
	in := positionLexemes(
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(")", token.R_PAREN),
	)

	exp := []tree.Node{
		tree.NumLit{Snippet: in[1].Snippet, Val: number.New("1")},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_ParenExpr_2(t *testing.T) {

	// (1 + 2)
	in := positionLexemes(
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("+", token.ADD),
		token.MakeTok("2", token.NUMBER),
		token.MakeTok(")", token.R_PAREN),
	)

	exp := []tree.Node{
		tree.BinaryExpr{
			Snippet: superSnip(in[1], in[3]),
			Left:    tree.NumLit{Snippet: in[1].Snippet, Val: number.New("1")},
			Op:      in[2].Token,
			OpPos:   in[2].Snippet,
			Right:   tree.NumLit{Snippet: in[3].Snippet, Val: number.New("2")},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_ParenExpr_3(t *testing.T) {

	// ((1 + 2))
	in := positionLexemes(
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("+", token.ADD),
		token.MakeTok("2", token.NUMBER),
		token.MakeTok(")", token.R_PAREN),
		token.MakeTok(")", token.R_PAREN),
	)

	exp := []tree.Node{
		tree.BinaryExpr{
			Snippet: superSnip(in[2], in[4]),
			Left:    tree.NumLit{Snippet: in[2].Snippet, Val: number.New("1")},
			Op:      in[3].Token,
			OpPos:   in[3].Snippet,
			Right:   tree.NumLit{Snippet: in[4].Snippet, Val: number.New("2")},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_ParenExpr_4(t *testing.T) {

	// ((1 + 2) * x) - y
	in := positionLexemes(
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("+", token.ADD),
		token.MakeTok("2", token.NUMBER),
		token.MakeTok(")", token.R_PAREN),
		token.MakeTok("*", token.MUL),
		token.MakeTok("x", token.IDENT),
		token.MakeTok(")", token.R_PAREN),
		token.MakeTok("-", token.SUB),
		token.MakeTok("y", token.IDENT),
	)

	// (1 + 2)
	add := tree.BinaryExpr{
		Snippet: superSnip(in[2], in[4]),
		Left:    tree.NumLit{Snippet: in[2].Snippet, Val: number.New("1")},
		Op:      in[3].Token,
		OpPos:   in[3].Snippet,
		Right:   tree.NumLit{Snippet: in[4].Snippet, Val: number.New("2")},
	}

	// ((1 + 2) * x)
	mul := tree.BinaryExpr{
		Snippet: superSnip(in[2], in[7]),
		Left:    add,
		Op:      in[6].Token,
		OpPos:   in[6].Snippet,
		Right:   tree.Ident{Snippet: in[7].Snippet, Val: "x"},
	}

	exp := []tree.Node{
		// ((1 + 2) * x) - y
		tree.BinaryExpr{
			Snippet: superSnip(in[2], in[10]),
			Left:    mul,
			Op:      in[9].Token,
			OpPos:   in[9].Snippet,
			Right:   tree.Ident{Snippet: in[10].Snippet, Val: "y"},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_SpellCall_1(t *testing.T) {

	// @Print()
	in := positionLexemes(
		token.MakeTok("@Print", token.SPELL),
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok(")", token.R_PAREN),
	)

	exp := []tree.Node{
		tree.SpellCall{
			Snippet:  superSnip(in[0], in[2]),
			Name:     "Print",
			ArgCount: 0,
			Args:     []tree.Expr{},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_SpellCall_2(t *testing.T) {

	// @Print(1 + 2)
	in := positionLexemes(
		token.MakeTok("@Print", token.SPELL),
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("+", token.ADD),
		token.MakeTok("2", token.NUMBER),
		token.MakeTok(")", token.R_PAREN),
	)

	// (1 + 2)
	add := tree.BinaryExpr{
		Snippet: superSnip(in[2], in[4]),
		Left:    tree.NumLit{Snippet: in[2].Snippet, Val: number.New("1")},
		Op:      in[3].Token,
		OpPos:   in[3].Snippet,
		Right:   tree.NumLit{Snippet: in[4].Snippet, Val: number.New("2")},
	}

	exp := []tree.Node{
		tree.SpellCall{
			Snippet:  superSnip(in[0], in[5]),
			Name:     "Print",
			ArgCount: 1,
			Args:     []tree.Expr{add},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_SpellCall_3(t *testing.T) {

	// @Print(true, 1, "abc")
	in := positionLexemes(
		token.MakeTok("@Print", token.SPELL),
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("true", token.TRUE),
		token.MakeTok(",", token.DELIM),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok(",", token.DELIM),
		token.MakeTok(`"abc"`, token.STRING),
		token.MakeTok(")", token.R_PAREN),
	)

	exp := []tree.Node{
		tree.SpellCall{
			Snippet:  superSnip(in[0], in[7]),
			Name:     "Print",
			ArgCount: 3,
			Args: []tree.Expr{
				tree.BoolLit{Snippet: in[2].Snippet, Val: true},
				tree.NumLit{Snippet: in[4].Snippet, Val: number.New("1")},
				tree.StrLit{Snippet: in[6].Snippet, Val: `"abc"`},
			},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_SpellCall_4(t *testing.T) {

	// @Print((x))
	in := positionLexemes(
		token.MakeTok("@Print", token.SPELL),
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok("x", token.IDENT),
		token.MakeTok(")", token.R_PAREN),
		token.MakeTok(")", token.R_PAREN),
	)

	exp := []tree.Node{
		tree.SpellCall{
			Snippet:  superSnip(in[0], in[5]),
			Name:     "Print",
			ArgCount: 1,
			Args: []tree.Expr{
				tree.Ident{Snippet: in[3].Snippet, Val: "x"},
			},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}
