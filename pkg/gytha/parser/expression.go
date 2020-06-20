package parser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func parseExpressions(p *pipe) []Expression {
	// pattern := [expression {DELIM expression}] [TERMINATOR]

	var exps []Expression
	exp := parseExpression(p)

	for exp != nil {
		exps = append(exps, exp)

		if p.accept(TK_DELIMITER) {
			exp = parseExpression(p)

			if exp == nil {
				err.Panic(
					errMsg("parseExpressions", `expression`, p.past()),
					err.After(p.past()),
				)
			}

		} else {
			exp = nil
		}
	}

	p.accept(TK_TERMINATOR) // In some cases
	return exps
}

func parseExpression(p *pipe) Expression {

	// pattern := [term {operator term}]
	// term := [negation] (func_call | list_access | literal | group | list)

	var left Expression

	switch {
	case isNegation(p):
		left = parseNegation(p)
		return parseOperation(p, left, 0)

	case isSpellCall(p):
		left = parseSpellCall(p)
		return parseOperation(p, left, 0)

	case isFuncCall(p):
		left = parseFuncCall(p)
		return parseOperation(p, left, 0)

	case isListAccess(p):
		left = parseListAccess(p)
		return parseOperation(p, left, 0)

	case isLiteral(p):
		left = parseLiteral(p)
		return parseOperation(p, left, 0)

	case isGroup(p):
		left = parseGroup(p)
		return parseOperation(p, left, 0)

	case isList(p):
		return parseList(p)
	}

	return nil
}
