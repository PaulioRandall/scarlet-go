package group

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/stat"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

func group(clt *collector, gp *grp) error {

	for clt.matchGen(GE_TERMINATOR) {
		clt.next() // Ignore empty statements
	}

	var e error

	switch {
	case clt.matchGen(GE_SPELL):
		e = spell(clt, gp)

	default:
		msg := fmt.Sprintf("Expected statement, got %s", clt.buff.String())
		e = perror.NewBySnippet(msg, clt.buff)
	}

	if e != nil {
		return e
	}

	if clt.empty() {
		return nil
	}

	_, e = clt.expectGen(GE_TERMINATOR)
	return e
}

func spell(clt *collector, gp *grp) error {

	gp.st = ST_SPELL_CALL

	if e := clt.expectAppendGen(gp, GE_SPELL); e != nil {
		return e
	}

	if e := clt.expectAppendSub(gp, SU_PAREN_OPEN); e != nil {
		return e
	}

	if clt.notMatchSub(SU_PAREN_CLOSE) {
		if e := expressions(clt, gp); e != nil {
			return e
		}
	}

	if e := clt.expectAppendSub(gp, SU_PAREN_CLOSE); e != nil {
		return e
	}

	return nil
}

func expressions(clt *collector, gp *grp) error {

	for {
		if e := expression(clt, gp); e != nil {
			return e
		}

		if !clt.acceptAppendSub(gp, SU_VALUE_DELIM) {
			break
		}
	}

	return nil
}

func expression(clt *collector, gp *grp) error {

	switch {
	case clt.acceptAppendGen(gp, GE_LITERAL):
	case clt.acceptAppendGen(gp, GE_IDENTIFIER):

	default:
		msg := fmt.Sprintf("Expected expression, got %s", clt.buff.String())
		return perror.NewBySnippet(msg, clt.buff)
	}

	return nil
}
