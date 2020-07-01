package group

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"

	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

func accept_append_gen(clt *collector, gp *grp, ge GenType) bool {

	if clt.matchGen(ge) {
		gp.tks = append(gp.tks, clt.next())
		return true
	}

	return false
}

func expect_append_gen(clt *collector, gp *grp, ge GenType) error {

	tk, e := clt.expectGen(ge)
	if e != nil {
		return e
	}

	gp.tks = append(gp.tks, tk)
	return nil
}

func accept_append_sub(clt *collector, gp *grp, su SubType) bool {

	if clt.matchSub(su) {
		gp.tks = append(gp.tks, clt.next())
		return true
	}

	return false
}

func expect_append_sub(clt *collector, gp *grp, su SubType) error {

	tk, e := clt.expectSub(su)
	if e != nil {
		return e
	}

	gp.tks = append(gp.tks, tk)
	return nil
}

func accept_discard_gen(clt *collector, ge GenType) bool {

	if clt.matchGen(ge) {
		clt.next()
		return true
	}

	return false
}

func expect_discard_gen(clt *collector, ge GenType) error {
	_, e := clt.expectGen(ge)
	return e
}

func nextGroup(clt *collector, gp *grp) error {

	for accept_discard_gen(clt, GE_TERMINATOR) {
		// Ignore empty statements
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

	expect_discard_gen(clt, GE_TERMINATOR)
	return nil
}

func spell(clt *collector, gp *grp) error {

	if e := expect_append_gen(clt, gp, GE_SPELL); e != nil {
		return e
	}

	if e := expect_append_sub(clt, gp, SU_PAREN_OPEN); e != nil {
		return e
	}

	if clt.notMatchSub(SU_PAREN_CLOSE) {
		if e := expressions(clt, gp); e != nil {
			return e
		}
	}

	if e := expect_append_sub(clt, gp, SU_PAREN_CLOSE); e != nil {
		return e
	}

	return nil
}

func expressions(clt *collector, gp *grp) error {

	for {
		if e := expression(clt, gp); e != nil {
			return e
		}

		if clt.notMatchSub(SU_VALUE_DELIM) {
			break
		}
	}

	return nil
}

func expression(clt *collector, gp *grp) error {

	switch {
	case accept_append_gen(clt, gp, GE_LITERAL):
	case accept_append_gen(clt, gp, GE_IDENTIFIER):

	default:
		msg := fmt.Sprintf("Expected expression, got %s", clt.buff.String())
		return perror.NewBySnippet(msg, clt.buff)
	}

	return nil
}
