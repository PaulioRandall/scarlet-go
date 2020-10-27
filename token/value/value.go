package value

import (
	"fmt"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/number"
)

type (
	// Value represents a value at runtime.
	Value interface {

		// Name returns the name of the type.
		Name() string

		// Comparable returns true if the 'other' value can be compared with the
		// receiver.
		Comparable(other Value) bool

		// Equal returns true if the 'other' value is equal to the receiver.
		Equal(other Value) bool

		// String returns the human readable representation of the value.
		String() string
	}

	// Ident represents an identifier. They are not available to language users
	// but are accessible within spells.
	Ident string
	Str   string
	Bool  bool
	Num   struct{ number.Number }
)

func (Ident) Name() string { return "ident" }
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

func (id Ident) Valid() Bool {
	for i, ru := range string(id) {
		if i == 0 && ru == '_' {
			return false
		}

		if !unicode.IsLetter(ru) && ru != '_' {
			return false
		}
	}
	return true
}

func (a Bool) And(b Bool) Bool { return Bool(bool(a) && bool(b)) }
func (a Bool) Or(b Bool) Bool  { return Bool(bool(a) || bool(b)) }
