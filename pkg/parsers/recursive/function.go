package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isFuncDef(p *pipe) bool {
	return p.match(FUNC)
}

func parseFuncDef(p *pipe) Expression {
	// pattern := FUNC params (statement | block)

	f := FuncDef{
		Key: p.expect(`parseFuncDef`, FUNC),
	}

	f.Inputs, f.Outputs = parseFuncParams(p)

	if isBlock(p) {
		f.Body = parseBlock(p)
		return f
	}

	if isGuard(p) || isMatch(p) || isLoop(p) {
		f.Body = parseStatBlock(p)
		return f
	}

	err.Panic(`Not a block or statement`, err.At(p.peek()))
	return nil
}

func isExprFuncDef(p *pipe) bool {
	return p.match(EXPR_FUNC)
}

func parseExprFuncDef(p *pipe) Expression {
	// pattern := EXPR_FUNC params expression

	f := ExprFuncDef{
		Key: p.expect(`parseExprFuncDef`, EXPR_FUNC),
	}

	var outputs []Token = nil
	f.Inputs, outputs = parseFuncParams(p)

	if outputs != nil {
		err.Panic(
			`Output variables not allowed in expression functions`,
			err.At(p.peek()),
		)
	}

	f.Expr = parseExpression(p)
	p.expect(`parseExprFuncDef`, TERMINATOR)
	return f
}

func parseFuncParams(p *pipe) (in, out []Token) {
	// pattern := PAREN_OPEN [ids] PAREN_CLOSE

	p.expect(`parseFuncParams`, PAREN_OPEN)

	if p.match(IDENTIFIER) || p.match(OUTPUT) {
		in, out = parseFuncParamIds(p)
	}

	p.expect(`parseFuncParams`, PAREN_CLOSE)

	return in, out
}

func parseFuncParamIds(p *pipe) (in []Token, out []Token) {
	// pattern := [^] ID {DELIM [^] ID}

	for {
		if p.accept(OUTPUT) {
			tk := p.expect(`parseFuncParamIds`, IDENTIFIER)
			out = append(out, tk)

		} else {
			tk := p.expect(`parseFuncParamIds`, IDENTIFIER)
			in = append(in, tk)
		}

		if !p.accept(DELIMITER) {
			break
		}
	}

	return
}

func isFuncCall(p *pipe) (is bool) {
	return p.matchSequence(IDENTIFIER, PAREN_OPEN)
}

func parseFuncCall(p *pipe) Expression {
	// pattern := ID PAREN_OPEN {expression} PAREN_CLOSE

	id := p.expect(`parseFuncCall`, IDENTIFIER)
	left := Identifier{id}

	p.expect(`parseFuncCall`, PAREN_OPEN)

	f := FuncCall{
		ID:     left,
		Inputs: parseExpressions(p),
	}

	p.expect(`parseFuncCall`, PAREN_CLOSE)
	return f
}
