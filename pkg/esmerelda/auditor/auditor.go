package auditor

// "Exactly what they were can't be described in normal language. Some people
// might call them cherubs, although there was nothing rosy-cheeked about them.
// They might be rumored among those who see to it that gravity operates and
// that time stays separate from space. Call them auditors. Auditors of
// reality."
// - 'Reaper Man' by Terry Pratchett

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type StatementIterator interface {
	Next() (Expression, error)
}

type errList struct {
	errs []error
}

func (el *errList) add(msg string, tk Token) {
	e := err.New(msg, err.At(tk))
	el.errs = append(el.errs, e)
}

func AuditStatements(si StatementIterator) []error {

	el := &errList{}

	for {

		st, e := si.Next()
		if e != nil {
			return []error{e}
		}

		if st == nil {
			return el.errs
		}

		statement(el, st)
	}

	return el.errs
}

func statements(el *errList, sts []Expression) {
	for _, st := range sts {
		statement(el, st)
	}
}

func statement(el *errList, st Expression) {

	switch st.Kind() {
	case ST_GUARD:
		v, _ := st.(Guard)
		guard(el, v)
	}
}

func guard(el *errList, g Guard) {
	expectBoolResult(el, g.Condition())
}

func expectBoolResult(el *errList, ex Expression) {

	if v, ok := ex.(Operation); ok {

		switch ty := v.Operator().Type(); {
		case IsComparisonType(ty):
			return

		case IsBoolLogicType(ty):
			expectBoolResult(el, v.Left())
			expectBoolResult(el, v.Right())
			return
		}

		el.add("Expected condition", v.Operator())
		return
	}

	if _, ok := ex.(Exists); ok {
		return
	}

	// TODO: function & spell calls
}
