package parser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

func isAssignment(p *pipe) bool {
	// match := DEF
	// match := VOID | (ID (DELIM | ASSIGN | GUARD_OPEN))

	return p.match(TK_DEFINITION) ||
		p.matchSequence(TK_IDENTIFIER, TK_DELIMITER) ||
		p.matchSequence(TK_IDENTIFIER, TK_ASSIGNMENT) ||
		p.matchSequence(TK_IDENTIFIER, TK_GUARD_OPEN) ||
		p.matchSequence(TK_VOID, TK_DELIMITER) ||
		p.matchSequence(TK_VOID, TK_ASSIGNMENT)
}

func parseAssignment(p *pipe) Assignment {
	// pattern := [DEF] assign_target {assign_target} ASSIGN expression {expression}

	a := Assignment{
		Fixed:   p.accept(TK_DEFINITION),
		Targets: parseAssignTargets(p),
		Assign:  p.expect(`parseAssignment`, TK_ASSIGNMENT),
		Exprs:   parseAssignExprs(p),
	}

	if a.Exprs == nil {
		err.Panic(
			errMsg("parseAssignment", `expression`, p.peek()),
			err.At(p.peek()),
		)
	}

	return a
}

func parseAssignExprs(p *pipe) []Expression {
	switch {
	case isFuncDef(p):
		return []Expression{parseFuncDef(p)}
	case isExprFuncDef(p):
		return []Expression{parseExprFuncDef(p)}
	}

	return parseExpressions(p)
}

func parseAssignTargets(p *pipe) []AssignTarget {
	// pattern := assignTarget { DELIM assignTarget }

	var ats []AssignTarget

	for !p.itr.Empty() {

		at := parseAssignTarget(p)
		ats = append(ats, at)

		if !p.accept(TK_DELIMITER) {
			break
		}
	}

	return ats
}

func parseAssignTarget(p *pipe) AssignTarget {
	// pattern := ID [GUARD_OPEN (NUMBER | ID) GUARD_CLOSE]

	at := AssignTarget{
		ID: p.expectOneOf(`parseAssignTarget`, TK_IDENTIFIER, TK_VOID),
	}

	if p.accept(TK_GUARD_OPEN) {

		switch {
		case p.matchAny(TK_LIST_START, TK_LIST_END):
			at.Index = ListItemRef{p.next()}

		case p.match(TK_NUMBER):
			at.Index = parseExpression(p)

		case p.match(TK_IDENTIFIER):
			at.Index = parseExpression(p)

		default:
			err.Panic(
				errMsg("parseAssignTarget", `NUMBER or ID`, p.peek()),
				err.At(p.peek()),
			)
		}

		p.expect(`parseAssignTarget`, TK_GUARD_CLOSE)
	}

	return at
}
