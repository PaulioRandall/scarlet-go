package value

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/number"
)

type Value interface {
	Name() string
	Comparable(other Value) bool
	Equal(other Value) bool
	String() string
}

type (
	Ident string
	Str   string
	Bool  bool
	Num   struct{ number.Number }
)

func (Ident) Name() string { return "string" }
func (Str) Name() string   { return "string" }
func (Bool) Name() string  { return "bool" }
func (Num) Name() string   { return "number" }

func (a Ident) Comparable(b Value) bool { return Str(a).Comparable(b) }
func (a Str) Comparable(b Value) bool   { _, ok := b.(Str); return ok }
func (a Bool) Comparable(b Value) bool  { _, ok := b.(Bool); return ok }
func (a Num) Comparable(b Value) bool   { _, ok := b.(Num); return ok }

func (a Ident) Equal(b Value) bool {
	return a.Comparable(b) && Str(a) == b.(Str)
}
func (a Str) Equal(b Value) bool {
	return a.Comparable(b) && a == b.(Str)
}
func (a Bool) Equal(b Value) bool {
	return a.Comparable(b) && a == b.(Bool)
}
func (a Num) Equal(b Value) bool {
	return a.Comparable(b) && a.Number.Equal(b.(Num).Number)
}

func (a Ident) String() string { return string(a) }
func (a Str) String() string   { return string(a) }
func (a Bool) String() string  { return fmt.Sprintf("%v", bool(a)) }
func (a Num) String() string   { return a.Number.String() }

func (a Bool) And(b Bool) Bool { return Bool(bool(a) && bool(b)) }
func (a Bool) Or(b Bool) Bool  { return Bool(bool(a) || bool(b)) }
