package check

import (
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/perror"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/shared/token/types"
)

func next(chk *checker) error {

	e := spell(chk)
	if e != nil {
		return e
	}

	return chk.expect(GE_TERMINATOR)
}

func spell(chk *checker) error {

	e := chk.expect(GE_SPELL)
	if e != nil {
		return e
	}

	e = chk.expect(SU_PAREN_OPEN)
	if e != nil {
		return e
	}

	if chk.accept(SU_PAREN_CLOSE) {
		return nil
	}

MORE:

	switch {
	case chk.accept(SU_IDENTIFIER):
	case chk.accept(GE_LITERAL):
	default:
		return perror.NewBySnippet("Unexpected token", chk.buff)
	}

	if chk.accept(SU_VALUE_DELIM) {
		goto MORE
	}

	return chk.expect(SU_PAREN_CLOSE)
}
