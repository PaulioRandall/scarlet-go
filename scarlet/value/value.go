package value

import (
	"fmt"
	"strconv"
	"unicode"
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

	Con interface {
		Value
		Len() int64
		CanBeKey(Value) bool
		CanHold(Value) bool
		Delete(Value) (Con, Value)
	}

	OrdCon interface {
		Con
		InRange(int64) bool
		At(int64) Value
		Index(Value) int64
		Slice(int64, int64) OrdCon
		PushFront(...Value) OrdCon
		PushBack(...Value) OrdCon
		PopFront() (OrdCon, Value)
		PopBack() (OrdCon, Value)
	}

	MutOrdCon interface {
		OrdCon
		Set(Value, Value) MutOrdCon
	}

	Ident string
	Bool  bool
	Num   float64
)

func (Ident) Name() string { return "ident" }
func (Bool) Name() string  { return "bool" }
func (Num) Name() string   { return "number" }

func (a Ident) Comparable(b Value) bool { return Str(a).Comparable(b) }
func (a Bool) Comparable(b Value) bool  { _, ok := b.(Bool); return ok }
func (a Num) Comparable(b Value) bool   { _, ok := b.(Num); return ok }

func (a Ident) Equal(b Value) bool {
	return a.Comparable(b) && Str(a) == b.(Str)
}
func (a Bool) Equal(b Value) bool {
	return a.Comparable(b) && a == b.(Bool)
}
func (a Num) Equal(b Value) bool {
	return a.Comparable(b) && a == b.(Num)
}

func (a Ident) String() string { return string(a) }
func (a Bool) String() string  { return fmt.Sprintf("%v", bool(a)) }
func (a Num) String() string {
	return strconv.FormatFloat(float64(a), 'f', -1, 64)
}

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
func (a Num) Int() int64       { return int64(a) }
func (a Ident) ToStr() Str     { return Str(string(a)) }
