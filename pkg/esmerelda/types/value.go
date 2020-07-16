package types

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/number"
)

type Value interface {
	Equal(other Value) bool
	Comparable(other Value) bool
	String() string
}

func BuiltinValueOf(val interface{}) Value {

	switch v := val.(type) {
	case int:
		return Int(v)

	case bool:
		return Bool(v)

	case number.Number:
		return Num{v}

	case string:
		return Str(v)
	}

	panic(fmt.Sprintf("No builtin type for value %v", val))
}
