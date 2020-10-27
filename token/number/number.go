package number

import (
	"fmt"
)

type Number interface {
	fmt.Stringer
	Copy() Number
	Integer() int64

	Inc(int64)
	Dec(int64)
	Neg()

	Add(Number)
	Sub(Number)
	Mul(Number)
	Div(Number)
	Mod(Number)

	Equal(Number) bool
	NotEqual(Number) bool
	Less(Number) bool
	LessOrEqual(Number) bool
	More(Number) bool
	MoreOrEqual(Number) bool
}

func New(numStr string) Number {
	return newShopspring(numStr)
}

func NewFromInt(i int64) Number {
	return newFromIntShopspring(i)
}

func NewFromFloat(f float64) Number {
	return newFromFloatShopspring(f)
}
