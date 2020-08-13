package types

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/shared/number"
)

type Delim struct{} // Internal use only

func (Delim) Name() string {
	return "delim"
}

func (Delim) Equal(Value) bool {
	return false
}

func (Delim) Comparable(Value) bool {
	return false
}

func (Delim) String() string {
	return ""
}

type Bool bool

func (Bool) Name() string {
	return "bool"
}

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

func (Num) Name() string {
	return "number"
}

func (a Num) Equal(b Value) bool {
	return a.Comparable(b) && a.Equal(b.(Num))
}

func (a Num) Comparable(b Value) bool {
	_, ok := b.(Num)
	return ok
}

func (a Num) String() string {
	return a.Number.String()
}

type Str string

func (Str) Name() string {
	return "string"
}

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
