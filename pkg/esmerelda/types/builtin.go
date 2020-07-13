package types

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/number"
)

func BuiltinValueOf(val interface{}) Value {

	switch v := val.(type) {
	case bool:
		return Bool(v)

	case number.Number:
		return Num{v}

	case string:
		return Str(v)
	}

	panic(fmt.Sprintf("No builtin type for value %v", val))
}

type Bool bool

func (a Bool) Equal(b Value) bool {
	return a.Comparable(b) && a == b.(Bool)
}

func (a Bool) Comparable(b Value) bool {
	_, ok := b.(Bool)
	return ok
}

func (a Bool) String() string {
	return fmt.Sprintf("%v", bool(a))
}

type Num struct {
	number.Number
}

func (a Num) Equal(b Value) bool {
	return a.Comparable(b) && a.Equal(b.(Num))
}

func (a Num) Comparable(b Value) bool {
	_, ok := b.(Num)
	return ok
}

func (a Num) String() string {
	return a.String()
}

type Str string

func (a Str) Equal(b Value) bool {
	return a.Comparable(b) && a == b.(Str)
}

func (a Str) Comparable(b Value) bool {
	_, ok := b.(Str)
	return ok
}

func (a Str) String() string {
	return string(a)
}
