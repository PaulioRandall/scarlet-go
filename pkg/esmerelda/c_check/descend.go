package check

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
)

func next(chk *checker) error {

	e := spell(chk)
	if e != nil {
		return e
	}

	return chk.expect(PR_TERMINATOR)
}

func spell(chk *checker) error {

	e := chk.expect(PR_SPELL)
	if e != nil {
		return e
	}

	e = chk.expect(PR_PARENTHESIS, PR_OPENER)
	if e != nil {
		return e
	}

	if chk.accept(PR_PARENTHESIS, PR_CLOSER) {
		return nil
	}

MORE:

	switch {
	case chk.accept(PR_TERM):
	default:
		return perror.NewBySnippet("Unexpected token", chk.buff)
	}

	if chk.accept(PR_SEPARATOR) {
		goto MORE
	}

	return chk.expect(PR_PARENTHESIS, PR_CLOSER)
}
