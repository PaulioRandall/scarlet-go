package gytha

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isFuncDef(p *pipe) bool {
	return p.match(TK_FUNCTION)
}

func parseFuncDef(p *pipe) Expression {
	// pattern := FUNC params (statement | block)

	f := FuncDef{
		Key: p.expect(`parseFuncDef`, TK_FUNCTION),
	}

	f.Inputs, f.Outputs = parseFuncParams(p)

	if isBlock(p) {
		f.Body = parseBlock(p)
		return f
	}

	if isGuard(p) || isWhen(p) {
		f.Body = parseStatBlock(p)
		return f
	}

	err.Panic(`Not a block or statement`, err.At(p.peek()))
	return nil
}

func isExprFuncDef(p *pipe) bool {
	return p.match(TK_EXPR_FUNC)
}

func parseExprFuncDef(p *pipe) Expression {
	// pattern := EXPR_FUNC params expression

	f := ExprFuncDef{
		Key: p.expect(`parseExprFuncDef`, TK_EXPR_FUNC),
	}

	var outputs []OutputParam
	f.Inputs, outputs = parseFuncParams(p)

	if outputs != nil {
		err.Panic(
			`Output variables not allowed in expression functions`,
			err.At(p.peek()),
		)
	}

	f.Expr = parseExpression(p)
	p.expect(`parseExprFuncDef`, TK_TERMINATOR)
	return f
}

func parseFuncParams(p *pipe) (in []Token, out []OutputParam) {
	// pattern := PAREN_OPEN [ids] PAREN_CLOSE

	p.expect(`parseFuncParams`, TK_PAREN_OPEN)

	if p.match(TK_IDENTIFIER) || p.match(TK_OUTPUT) {
		in, out = parseFuncParamIds(p)
	}

	p.expect(`parseFuncParams`, TK_PAREN_CLOSE)

	return in, out
}

func parseFuncParamIds(p *pipe) (in []Token, out []OutputParam) {
	// pattern := param {DELIM param}
	// param := (ID | (OUTPUT ID ASSIGN expression))

	for {
		if p.accept(TK_OUTPUT) {
			tk := parseOutputParam(p)
			out = append(out, tk)

		} else {
			tk := p.expect(`parseFuncParamIds`, TK_IDENTIFIER)
			in = append(in, tk)
		}

		if !p.accept(TK_DELIMITER) {
			break
		}
	}

	return
}

func parseOutputParam(p *pipe) OutputParam {
	// pattern := ID [ASSIGN expression]

	o := OutputParam{
		ID: Identifier{
			Tk: p.expect(`parseFuncParamIds`, TK_IDENTIFIER),
		},
	}

	if p.accept(TK_ASSIGNMENT) {
		o.Expr = parseExpression(p)

		if o.Expr == nil {
			err.Panic(
				fmt.Sprintf(`Missing expression after %s`, p.past()),
				err.At(p.past()),
			)
		}
	}

	return o
}

func isFuncCall(p *pipe) (is bool) {
	return p.matchSequence(TK_IDENTIFIER, TK_PAREN_OPEN)
}

func parseFuncCall(p *pipe) Expression {
	// pattern := ID PAREN_OPEN {expression} PAREN_CLOSE

	id := p.expect(`parseFuncCall`, TK_IDENTIFIER)
	left := Identifier{id}

	p.expect(`parseFuncCall`, TK_PAREN_OPEN)

	f := FuncCall{
		ID:     left,
		Inputs: parseExpressions(p),
	}

	p.expect(`parseFuncCall`, TK_PAREN_CLOSE)
	return f
}

func isSpellCall(p *pipe) bool {
	return p.match(TK_SPELL)
}

func parseSpellCall(p *pipe) Expression {

	p.expect(`parseSpell`, TK_SPELL)
	id := p.expect(`parseSpell`, TK_IDENTIFIER)

	p.expect(`parseSpell`, TK_PAREN_OPEN)
	inputs := parseExpressions(p)
	p.expect(`parseSpell`, TK_PAREN_CLOSE)

	return SpellCall{
		ID:     id,
		Inputs: inputs,
	}
}
