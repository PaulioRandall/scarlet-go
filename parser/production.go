package parser

import (
	"github.com/PaulioRandall/scarlet-go/parser/eval"
	"github.com/PaulioRandall/scarlet-go/token"
)

// ID_ARRAY         := ID_OR_VOID { "," ID_OR_VOID } .
func matchIdArray(tc *TokenCollector) ([]eval.Expr, ParseErr) {

	ev, _ := matchIdOrVoid(tc)

	if ev == nil {
		return nil, nil
	}

	ids := []eval.Expr{ev}
	matchMoreIds(ids, tc)

	return ids, nil
}

// ID_OR_VOID       := ID | "\_" .
func matchIdOrVoid(tc *TokenCollector) (eval.Expr, ParseErr) {

	t := tc.Read()

	if tc.Err() != nil {
		return nil, tc.Err()
	}

	if t.Kind != token.ID && t.Kind != token.VOID {
		tc.PutBack(1)
		return nil, nil
	}

	return eval.NewForID(t), nil
}

// ID_ARRAY         := ... { "," ID_OR_VOID } .
func matchMoreIds(ids []eval.Expr, tc *TokenCollector) ParseErr {
	for tc.Peek().Kind == token.ID_DELIM {

		// Skip the delimiter
		if _ = tc.Read(); tc.Err() != nil {
			return tc.Err()
		}

		ev, _ := matchIdOrVoid(tc)

		if ev == nil {
			return NewParseErr("Expected ID token", nil, tc.Peek())
		}

		ids = append(ids, ev)
	}

	return nil
}
