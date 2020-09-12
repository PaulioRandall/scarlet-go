package types

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/number"
)

type Value interface {
	Name() string
	Equal(other Value) bool
	Comparable(other Value) bool
	String() string
}

func BuiltinValueOf(val interface{}) Value {

	switch v := val.(type) {
	case nil:
	case bool:
		return Bool(v)

	case number.Number:
		return Num{v}

	case string:
		return Str(v)
	}

	panic(fmt.Sprintf("No builtin type for value %v", val))
}
