package check

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token/types"
)

func next(chk *checker) error {

	e := spell(chk)
	if e != nil {
		return e
	}

	return chk.expect(GEN_TERMINATOR)
}

func spell(chk *checker) error {

	e := chk.expect(GEN_SPELL)
	if e != nil {
		return e
	}

	e = chk.expect(SUB_PAREN_OPEN)
	if e != nil {
		return e
	}

	if chk.accept(SUB_PAREN_CLOSE) {
		return nil
	}

MORE:

	switch {
	case chk.accept(SUB_IDENTIFIER):
	case chk.accept(GEN_LITERAL):
	default:
		return perror.NewBySnippet("Unexpected token", chk.buff)
	}

	if chk.accept(SUB_VALUE_DELIM) {
		goto MORE
	}

	return chk.expect(SUB_PAREN_CLOSE)
}
