package z_recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
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
