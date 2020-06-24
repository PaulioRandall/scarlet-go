package runtime

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
)

func evalStatements(ctx *Context, stats []Expression) error {

	for _, st := range stats {
		e := evalStatement(ctx, st)
		if e != nil {
			return e
		}
	}

	return nil
}

func evalStatement(ctx *Context, st Expression) error {

	switch st.Kind() {
	case ST_ASSIGNMENT:
		return evalAssignment(ctx, st.(Assignment))
	}

	return err.NewBySnippet("Unknown statement type", st)
}

func evalAssignment(ctx *Context, as Assignment) error {
	return nil
}
