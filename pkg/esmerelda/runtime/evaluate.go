package runtime

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
)

func evalStatements(ctx *Context, stats []Expression) {
	for _, st := range stats {
		evalStatement(ctx, st)
	}
}

func evalStatement(ctx *Context, st Expression) {
	switch st.Kind() {
	}
}
