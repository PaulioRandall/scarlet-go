package recursive

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isFuncDef(p *pipe) bool {
	return p.match(token.FUNC)
}

// Expects the following token pattern:
// pattern := FUNC params (statement | block)
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
func parseFuncParams(p *pipe) (in, out []token.Token) {

	p.expect(`parseFuncParams`, token.PAREN_OPEN)

	if p.match(token.ID) {
		in = parseFuncParamIds(p)
	}

	if p.accept(token.RETURNS) {
		out = parseFuncParamIds(p)
	}

	p.expect(`parseFuncParams`, token.PAREN_CLOSE)

	return in, out
}

// Expects the following token pattern:
// pattern := ID {DELIM ID}
func parseFuncParamIds(p *pipe) []token.Token {

	var ids []token.Token

	for {
		tk := p.expect(`parseFuncParamIds`, token.ID)
		ids = append(ids, tk)

		if !p.accept(token.DELIM) {
			break
		}
	}

	return ids
}

func isFuncBlock(p *pipe) bool {
	return p.match(token.BLOCK_OPEN)
}

// Expects the following token pattern:
// pattern := BLOCK_OPEN {statement} BLOCK_CLOSE
func parseFuncBlock(p *pipe) st.Block {
	return st.Block{
		Open:  p.expect(`parseFuncBlock`, token.BLOCK_OPEN),
		Stats: parseStatements(p),
		Close: p.expect(`parseFuncBlock`, token.BLOCK_CLOSE),
	}
}

// Expects the following token pattern:
// pattern := statement
func parseFuncStatement(p *pipe) st.Block {
	return st.Block{
		Open:  p.peek(),
		Stats: []st.Statement{parseStatement(p)},
		Close: p.past(),
	}
}

func isFuncCall(p *pipe) (is bool) {
	return p.matchSequence(token.ID, token.PAREN_OPEN)
}

// Expects the following token pattern:
// pattern := ID PAREN_OPEN {expression} PAREN_CLOSE
func parseFuncCall(p *pipe) st.Expression {

	id := p.expect(`parseFuncCall`, token.ID)
	left := st.Identifier(id)

	p.expect(`parseFuncCall`, token.PAREN_OPEN)

	f := st.FuncCall{
		ID:    left,
		Input: parseExpressions(p),
	}

	p.expect(`parseFuncCall`, token.PAREN_CLOSE)
	return f
}
