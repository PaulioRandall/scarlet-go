package z_recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func parseExpressions(p *pipe) []Expression {
	// pattern := [expression {DELIM expression}] [TERMINATOR]

	var exps []Expression
	exp := parseExpression(p)

	for exp != nil {
		exps = append(exps, exp)

		if p.accept(DELIMITER) {
			exp = parseExpression(p)

			if exp == nil {
				panic(err("parseExpressions", p.past(), 1, `Expected expression`))
			}

		} else {
			exp = nil
		}
	}

	p.accept(TERMINATOR) // In some cases
	return exps
}

func parseExpression(p *pipe) Expression {
	// pattern := func_call | list_access | literal | group | list

	var left Expression

	switch {
	//case isFuncCall(p):
	//	return parseFuncCall(p)

	//case isListAccess(p):
	//	return parseListAccess(p)

	case isLiteral(p):
		left = parseLiteral(p)
		return parseOperation(p, left, 0)

	case isGroup(p):
		left = parseGroup(p)
		return parseOperation(p, left, 0)

		//case isList(p):
		//		return parseList(p)
	}

	return nil
}
