package z_recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func parseOperation(p *pipe, left Expression, leftPriority int) Expression {
	// pattern := {operator operand}
	// operand := literal | expression

	// Warning: this is where parsing gets a little complicated!!

	op := p.peek()
	m := op.Morpheme()

	if leftPriority >= m.Precedence() {
		// Any token that is not an operator has a precedence of zero, so the left
		// hand expression will always be returned in such a case.
		return left
	}

	// Because operator not taken yet.
	p.expect(`parseOperation`, m)

	// Parse the terminal or expression on the right.
	right := parseSubOperation(p)

	// Recursively parse the expression on the right until an operator with
	// precedence less or equal to this one is encountered.
	right = parseOperation(p, right, m.Precedence())

	left = Operation{left, op, right}

	// Parse the remaining operations in this expression.
	left = parseOperation(p, left, leftPriority)

	return left
}

func parseSubOperation(p *pipe) Expression {
	// pattern := func_call | literal | group

	switch {
	//	case isFuncCall(p):
	//		return parseFuncCall(p)

	case isLiteral(p):
		return parseLiteral(p)

	case isGroup(p):
		return parseGroup(p)
	}

	panic(unexpected("parseRightSide", p.peek(), `function_call | literal | group`))
}

func isLiteral(p *pipe) bool {
	return p.matchAny(
		IDENTIFIER, // Yes I know, need to sort it out
		VOID,
		BOOL,
		NUMBER,
		STRING,
		TEMPLATE,
	)
}

func parseLiteral(p *pipe) Expression {
	tk := p.next()

	if tk.Morpheme() == IDENTIFIER {
		return Identifier{tk}
	} else {
		return Value{tk}
	}
}

func isGroup(p *pipe) bool {
	return p.match(PAREN_OPEN)
}

func parseGroup(p *pipe) Expression {
	// pattern := PAREN_OPEN expression PAREN_CLOSE

	p.expect(`group`, PAREN_OPEN)

	g := parseExpression(p)
	if g == nil {
		panic(unexpected("group", p.past(), ANOTHER.String()))
	}

	p.expect(`group`, PAREN_CLOSE)
	return g
}
