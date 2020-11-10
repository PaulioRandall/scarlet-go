package parser

import (
	"testing"

	"github.com/PaulioRandall/scarlet-go/scarlet/token"
	"github.com/PaulioRandall/scarlet-go/scarlet/tree"

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
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.SingleAssign{
			Left:  tree.Ident{Val: "x"},
			Right: tree.NumLit{Val: float64(1)},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.SingleAssign{
			Left:  tree.Ident{Val: "x"},
			Right: tree.Ident{Val: "y"},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_SingleAssign_3(t *testing.T) {

	// _ := 1
	in := positionLexemes(
		token.MakeTok("_", token.VOID),
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.SingleAssign{
			Left:  tree.AnonIdent{},
			Right: tree.NumLit{Val: float64(1)},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_MultiAssign_1(t *testing.T) {

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
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.MultiAssign{
			Left: []tree.Assignee{
				tree.Ident{Val: "x"},
				tree.Ident{Val: "y"},
				tree.Ident{Val: "z"},
			},
			Right: []tree.Expr{
				tree.BoolLit{Val: true},
				tree.NumLit{Val: float64(1)},
				tree.StrLit{Val: `"text"`},
			},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_MultiAssign_2(t *testing.T) {

	// _, y, _ := true, 1, "abc"
	in := positionLexemes(
		token.MakeTok("_", token.VOID), // 0
		token.MakeTok(",", token.DELIM),
		token.MakeTok("y", token.IDENT), // 2
		token.MakeTok(",", token.DELIM),
		token.MakeTok("_", token.VOID), // 4
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("true", token.TRUE), // 6
		token.MakeTok(",", token.DELIM),
		token.MakeTok("1", token.NUMBER), // 8
		token.MakeTok(",", token.DELIM),
		token.MakeTok(`"text"`, token.STRING), // 10
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.MultiAssign{
			Left: []tree.Assignee{
				tree.AnonIdent{},
				tree.Ident{Val: "y"},
				tree.AnonIdent{},
			},
			Right: []tree.Expr{
				tree.BoolLit{Val: true},
				tree.NumLit{Val: float64(1)},
				tree.StrLit{Val: `"text"`},
			},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_AsymAssign_1(t *testing.T) {

	// x, y := @Print()
	in := positionLexemes(
		token.MakeTok("x", token.IDENT), // 0
		token.MakeTok(",", token.DELIM),
		token.MakeTok("y", token.IDENT), // 2
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("@Print", token.SPELL), // 4
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok(")", token.R_PAREN), // 6
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.AsymAssign{
			Left: []tree.Assignee{
				tree.Ident{Val: "x"},
				tree.Ident{Val: "y"},
			},
			Right: tree.SpellCall{
				Name: "Print",
				Args: []tree.Expr{},
			},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_AsymAssign_2(t *testing.T) {

	// _, _ := @Print()
	in := positionLexemes(
		token.MakeTok("_", token.VOID), // 0
		token.MakeTok(",", token.DELIM),
		token.MakeTok("_", token.VOID), // 2
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("@Print", token.SPELL), // 4
		token.MakeTok("(", token.L_PAREN),
		token.MakeTok(")", token.R_PAREN), // 6
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.AsymAssign{
			Left: []tree.Assignee{
				tree.AnonIdent{},
				tree.AnonIdent{},
			},
			Right: tree.SpellCall{
				Name: "Print",
				Args: []tree.Expr{},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.Ident{Val: "x"},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.BinaryExpr{
			Left:  tree.NumLit{Val: float64(1)},
			Op:    tokenToOperator(in[1].Token),
			Right: tree.NumLit{Val: float64(2)},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.BinaryExpr{
			Left:  tree.BoolLit{Val: true},
			Op:    tokenToOperator(in[1].Token),
			Right: tree.BoolLit{Val: false},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	add := tree.BinaryExpr{
		Left:  tree.NumLit{Val: float64(1)},
		Op:    tokenToOperator(in[1].Token),
		Right: tree.NumLit{Val: float64(2)},
	}

	exp := []tree.Node{
		tree.BinaryExpr{
			Left:  add,
			Op:    tokenToOperator(in[3].Token),
			Right: tree.NumLit{Val: float64(3)},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	mul := tree.BinaryExpr{
		Left:  tree.NumLit{Val: float64(2)},
		Op:    tokenToOperator(in[3].Token),
		Right: tree.NumLit{Val: float64(3)},
	}

	exp := []tree.Node{
		tree.BinaryExpr{
			Left:  tree.NumLit{Val: float64(1)},
			Op:    tokenToOperator(in[1].Token),
			Right: mul,
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
		token.MakeTok("\n", token.NEWLINE),
	)

	// 2 * 3
	mul := tree.BinaryExpr{
		Left:  tree.NumLit{Val: float64(2)},
		Op:    tokenToOperator(in[3].Token),
		Right: tree.NumLit{Val: float64(3)},
	}

	// 1 + (2 * 3)
	add := tree.BinaryExpr{
		Left:  tree.NumLit{Val: float64(1)},
		Op:    tokenToOperator(in[1].Token),
		Right: mul,
	}

	// 4 / 5
	div := tree.BinaryExpr{
		Left:  tree.NumLit{Val: float64(4)},
		Op:    tokenToOperator(in[7].Token),
		Right: tree.NumLit{Val: float64(5)},
	}

	// (4 / 5) % 6
	rem := tree.BinaryExpr{
		Left:  div,
		Op:    tokenToOperator(in[9].Token),
		Right: tree.NumLit{Val: float64(6)},
	}

	exp := []tree.Node{
		// (1 + 2 * 3) - (4 / 5 % 6)
		tree.BinaryExpr{
			Left:  add,
			Op:    tokenToOperator(in[5].Token),
			Right: rem,
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
		token.MakeTok("\n", token.NEWLINE),
	)

	right := tree.BinaryExpr{
		Left:  tree.NumLit{Val: float64(1)},
		Op:    tokenToOperator(in[3].Token),
		Right: tree.NumLit{Val: float64(2)},
	}

	exp := []tree.Node{
		tree.SingleAssign{
			Left:  tree.Ident{Val: "x"},
			Right: right,
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
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.NumLit{Val: float64(1)},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.BinaryExpr{
			Left:  tree.NumLit{Val: float64(1)},
			Op:    tokenToOperator(in[2].Token),
			Right: tree.NumLit{Val: float64(2)},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.BinaryExpr{
			Left:  tree.NumLit{Val: float64(1)},
			Op:    tokenToOperator(in[3].Token),
			Right: tree.NumLit{Val: float64(2)},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	// (1 + 2)
	add := tree.BinaryExpr{
		Left:  tree.NumLit{Val: float64(1)},
		Op:    tokenToOperator(in[3].Token),
		Right: tree.NumLit{Val: float64(2)},
	}

	// ((1 + 2) * x)
	mul := tree.BinaryExpr{
		Left:  add,
		Op:    tokenToOperator(in[6].Token),
		Right: tree.Ident{Val: "x"},
	}

	exp := []tree.Node{
		// ((1 + 2) * x) - y
		tree.BinaryExpr{
			Left:  mul,
			Op:    tokenToOperator(in[9].Token),
			Right: tree.Ident{Val: "y"},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.SpellCall{
			Name: "Print",
			Args: []tree.Expr{},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	// (1 + 2)
	add := tree.BinaryExpr{
		Left:  tree.NumLit{Val: float64(1)},
		Op:    tokenToOperator(in[3].Token),
		Right: tree.NumLit{Val: float64(2)},
	}

	exp := []tree.Node{
		tree.SpellCall{
			Name: "Print",
			Args: []tree.Expr{add},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.SpellCall{
			Name: "Print",
			Args: []tree.Expr{
				tree.BoolLit{Val: true},
				tree.NumLit{Val: float64(1)},
				tree.StrLit{Val: `"abc"`},
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
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.SpellCall{
			Name: "Print",
			Args: []tree.Expr{
				tree.Ident{Val: "x"},
			},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_Exist_1(t *testing.T) {

	// x?
	in := positionLexemes(
		token.MakeTok("x", token.IDENT),
		token.MakeTok("?", token.EXIST),
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.UnaryExpr{
			Term: tree.Ident{Val: "x"},
			Op:   tree.OP_EXIST,
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_Block_1(t *testing.T) {

	// {}
	in := positionLexemes(
		token.MakeTok("{", token.L_CURLY),
		token.MakeTok("}", token.R_CURLY),
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.Block{
			Stmts: []tree.Node{},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_Block_2(t *testing.T) {

	// {}
	in := positionLexemes(
		token.MakeTok("{", token.L_CURLY),
		token.MakeTok("x", token.IDENT),
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("1", token.NUMBER),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("y", token.IDENT),
		token.MakeTok(":=", token.ASSIGN),
		token.MakeTok("2", token.NUMBER),
		token.MakeTok("\n", token.NEWLINE),
		token.MakeTok("}", token.R_CURLY),
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.Block{
			Stmts: []tree.Node{
				tree.SingleAssign{
					Left:  tree.Ident{Val: "x"},
					Right: tree.NumLit{Val: float64(1)},
				},
				tree.SingleAssign{
					Left:  tree.Ident{Val: "y"},
					Right: tree.NumLit{Val: float64(2)},
				},
			},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}

func TestParse_Guard_1(t *testing.T) {

	// [true] {}
	in := positionLexemes(
		token.MakeTok("[", token.L_SQUARE),
		token.MakeTok("true", token.TRUE),
		token.MakeTok("]", token.R_SQUARE),
		token.MakeTok("{", token.L_CURLY),
		token.MakeTok("}", token.R_CURLY),
		token.MakeTok("\n", token.NEWLINE),
	)

	exp := []tree.Node{
		tree.Guard{
			Cond: tree.BoolLit{Val: true},
			Body: tree.Block{
				Stmts: []tree.Node{},
			},
		},
	}

	act, e := ParseAll(in)
	require.Nil(t, e, "ERROR: %+v", e)
	requireNodes(t, exp, act)
}
