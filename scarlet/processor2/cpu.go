package processor2

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

type (
	CPU interface {
		ADD(l, r value.Num) value.Num
		SUB(l, r value.Num) value.Num
		MUL(l, r value.Num) value.Num
		DIV(l, r value.Num) value.Num
		REM(l, r value.Num) value.Num

		AND(l, r value.Bool) value.Bool
		OR(l, r value.Bool) value.Bool

		LT(l, r value.Num) value.Num
		MT(l, r value.Num) value.Num
		LTE(l, r value.Num) value.Num
		MTE(l, r value.Num) value.Num

		EQU(l, r value.Value) value.Bool
		NEQ(l, r value.Value) value.Bool
	}

	SimpleCPU struct {
	}
)

func (SimpleCPU) ADD(l, r value.Num) value.Num {
	l.Number = l.Number.Copy()
	l.Number.Add(r.Number)
	return l
}

func (SimpleCPU) SUB(l, r value.Num) value.Num {
	l.Number = l.Number.Copy()
	l.Number.Sub(r.Number)
	return l
}

func (SimpleCPU) MUL(l, r value.Num) value.Num {
	l.Number = l.Number.Copy()
	l.Number.Mul(r.Number)
	return l
}

func (SimpleCPU) DIV(l, r value.Num) value.Num {
	l.Number = l.Number.Copy()
	l.Number.Div(r.Number)
	return l
}

func (SimpleCPU) REM(l, r value.Num) value.Num {
	l.Number = l.Number.Copy()
	l.Number.Div(r.Number)
	return l
}

func (SimpleCPU) AND(l, r value.Bool) value.Bool {
	return l && r
}

func (SimpleCPU) OR(l, r value.Bool) value.Bool {
	return l || r
}

func (SimpleCPU) LT(l, r value.Num) value.Num {
	l.Number = l.Number.Copy()
	l.Number.Less(r.Number)
	return l
}

func (SimpleCPU) MT(l, r value.Num) value.Num {
	l.Number = l.Number.Copy()
	l.Number.More(r.Number)
	return l
}

func (SimpleCPU) LTE(l, r value.Num) value.Num {
	l.Number = l.Number.Copy()
	l.Number.LessOrEqual(r.Number)
	return l
}

func (SimpleCPU) MTE(l, r value.Num) value.Num {
	l.Number = l.Number.Copy()
	l.Number.MoreOrEqual(r.Number)
	return l
}

func (SimpleCPU) EQU(l, r value.Value) value.Bool {
	return value.Bool(l.Equal(r))
}

func (SimpleCPU) NEQ(l, r value.Value) value.Bool {
	return value.Bool(!l.Equal(r))
}
