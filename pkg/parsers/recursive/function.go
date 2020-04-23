package recursive

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isFuncDef(p *pipe) bool {
	return p.inspect(token.FUNC)
}

// Expects the following token pattern:
// pattern := FUNC params (statement | block)
// block := BLOCK_OPEN {statement} BLOCK_CLOSE
// params := PAREN_OPEN [ids] [RETURNS ids] PAREN_CLOSE
// ids := ID {DELIM ID}
func parseFuncDef(p *pipe) st.Expression {

	f := st.FuncDef{
		Open: p.expect(`parseFuncDef`, token.FUNC),
	}

	f.Input, f.Output = parseFuncParams(p)

	if isFuncBlock(p) {
		f.Body = parseFuncBlock(p)
	} else {
		f.Body = parseFuncStatement(p)
	}

	return f
}

// Expects the following token pattern:
// pattern := PAREN_OPEN [ids] [RETURNS ids] PAREN_CLOSE
// ids := ID {DELIM ID}
func parseFuncParams(p *pipe) (in, out []token.Token) {

	p.expect(`parseFuncParams`, token.PAREN_OPEN)

	if p.inspect(token.ID) {
		in = parseFuncIds(p)
	}

	if p.accept(token.RETURNS) {
		out = parseFuncIds(p)
	}

	p.expect(`parseFuncParams`, token.PAREN_CLOSE)

	return in, out
}

// Expects the following token pattern:
// pattern := ID {DELIM ID}
func parseFuncIds(p *pipe) []token.Token {

	var ids []token.Token

	for {
		tk := p.expect(`parseFuncIds`, token.ID)
		ids = append(ids, tk)

		if !p.accept(token.DELIM) {
			break
		}
	}

	return ids
}

func isFuncBlock(p *pipe) bool {
	return p.inspect(token.BLOCK_OPEN)
}

// Expects the following token pattern:
// pattern := BLOCK_OPEN {statement} BLOCK_CLOSE
func parseFuncBlock(p *pipe) st.Block {
	return st.Block{
		Open:  p.expect(`parseFuncBlock`, token.BLOCK_OPEN),
		Stats: statements(p),
		Close: p.expect(`parseFuncBlock`, token.BLOCK_CLOSE),
	}
}

// Expects the following token pattern:
// pattern := statement
func parseFuncStatement(p *pipe) st.Block {
	return st.Block{
		Open:  p.snoop(),
		Stats: []st.Statement{statement(p)},
		Close: p.prior(),
	}
}

func isFuncCall(p *pipe) (is bool) {
	return p.isSequence(token.ID, token.PAREN_OPEN)
}

// Expects the following token pattern:
// pattern := ID PAREN_OPEN {expression} PAREN_CLOSE
func parseFuncCall(p *pipe) st.Expression {

	left := st.Identifier{
		Fixed:  false,
		Source: p.expect(`parseFuncCall`, token.ID),
	}

	p.expect(`parseFuncCall`, token.PAREN_OPEN)

	f := st.FuncCall{
		ID:    left,
		Input: parseExpressions(p),
	}

	p.expect(`parseFuncCall`, token.PAREN_CLOSE)
	return f
}
