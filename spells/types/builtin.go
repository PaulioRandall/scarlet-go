package types

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/number"
)

type Nil struct{} // Internal use only

func (Nil) Name() string {
	return "nil"
}

func (n Nil) Equal(o Value) bool {
	return n.Comparable(o)
}

func (n Nil) Comparable(o Value) bool {
	_, ok := o.(Nil)
	return ok
}

func (Nil) String() string {
	return ""
}

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

func (a Bool) And(b Bool) Bool {
	return Bool(bool(a) && bool(b))
}

func (a Bool) Or(b Bool) Bool {
	return Bool(bool(a) || bool(b))
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
	return a.Comparable(b) && a.Number.Equal(b.(Num).Number)
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
