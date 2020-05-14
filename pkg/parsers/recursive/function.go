package recursive

import (
	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isFuncDef(p *pipe) bool {
	return p.match(token.FUNC)
}

func parseFuncDef(p *pipe) st.Expression {
	// pattern := FUNC params (statement | block)

	f := st.FuncDef{
		Key: p.expect(`parseFuncDef`, token.FUNC),
	}

	f.Inputs, f.Outputs = parseFuncParams(p)

	if isBlock(p) {
		f.Body = parseBlock(p)
	} else {

		/*
			if len(f.Output) != 1 {
				panic(err("parseFuncDef", p.peek(),
				"Inline function bodies must have a single output parameter"))
			}
		*/
		f.Body = parseStatBlock(p)
	}

	return f
}

func parseFuncParams(p *pipe) (in, out []token.Token) {
	// pattern := PAREN_OPEN [ids] PAREN_CLOSE

	p.expect(`parseFuncParams`, token.PAREN_OPEN)

	if p.match(token.ID) || p.match(token.OUTPUT) {
		in, out = parseFuncParamIds(p)
	}

	p.expect(`parseFuncParams`, token.PAREN_CLOSE)

	return in, out
}

func parseFuncParamIds(p *pipe) (in []token.Token, out []token.Token) {
	// pattern := [^] ID {DELIM [^] ID}

	for {
		if p.accept(token.OUTPUT) {
			tk := p.expect(`parseFuncParamIds`, token.ID)
			out = append(out, tk)

		} else {
			tk := p.expect(`parseFuncParamIds`, token.ID)
			in = append(in, tk)
		}

		if !p.accept(token.DELIM) {
			break
		}
	}

	return
}

func isFuncCall(p *pipe) (is bool) {
	return p.matchSequence(token.ID, token.PAREN_OPEN)
}

func parseFuncCall(p *pipe) st.Expression {
	// pattern := ID PAREN_OPEN {expression} PAREN_CLOSE

	id := p.expect(`parseFuncCall`, token.ID)
	left := st.Identifier(id)

	p.expect(`parseFuncCall`, token.PAREN_OPEN)

	f := st.FuncCall{
		ID:     left,
		Inputs: parseExpressions(p),
	}

	p.expect(`parseFuncCall`, token.PAREN_CLOSE)
	return f
}
