package number

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
)

func shMake(num interface{}, exp int32) Number {

	switch v := num.(type) {
	case int:
		return &number{d: decimal.New(int64(v), exp)}

	case int32:
		return &number{d: decimal.New(int64(v), exp)}

	case int64:
		return &number{d: decimal.New(int64(v), exp)}

	case float32:
		return &number{d: decimal.NewFromFloatWithExponent(float64(v), exp)}

	case float64:
		return &number{d: decimal.NewFromFloatWithExponent(float64(v), exp)}
	}

	panic("SANITY CHECK! Unknown type")
}

func requireEqual(t *testing.T, expect, actual Number) {
	exp, _ := expect.(*number)
	act, _ := actual.(*number)
	exp.d, act.d = decimal.RescalePair(exp.d, act.d)
	require.Equal(t, exp, act)
}

func Test_SH1_1(t *testing.T) {
	// number.Integer returns the correct value

	exp := int64(1)
	act := shMake(1.1, 0).Integer()

	require.Equal(t, exp, act)
}

func Test_SH2_2(t *testing.T) {
	// number.Inc increments the number

	act := newShopspring("1")

	act.Inc(1)
	requireEqual(t, shMake(2, 0), act)

	act.Inc(3)
	requireEqual(t, shMake(5, 0), act)
}

func Test_SH3_1(t *testing.T) {
	// number.Dec decrements the number

	act := newShopspring("5")

	act.Dec(1)
	requireEqual(t, shMake(4, 0), act)

	act.Dec(3)
	requireEqual(t, shMake(1, 0), act)
}

func Test_SH4_1(t *testing.T) {
	// number.Add adds a specified amount to the number

	act := newShopspring("1")

	act.Add(newShopspring("1"))
	requireEqual(t, shMake(2, 0), act)

	act.Add(newShopspring("3"))
	requireEqual(t, shMake(5, 0), act)

	act.Add(newShopspring("-2"))
	requireEqual(t, shMake(3, 0), act)
}

func Test_SH5_1(t *testing.T) {
	// number.Sub subtracts a specified amount to the number

	act := newShopspring("5")

	act.Sub(newShopspring("1"))
	requireEqual(t, shMake(4, 0), act)

	act.Sub(newShopspring("3"))
	requireEqual(t, shMake(1, 0), act)

	act.Sub(newShopspring("-2"))
	requireEqual(t, shMake(3, 0), act)
}

func Test_SH6_1(t *testing.T) {
	// number.Mul multiplies a specified amount to the number

	act := newShopspring("2")

	act.Mul(newShopspring("2"))
	requireEqual(t, shMake(4, 0), act)

	act.Mul(newShopspring("-3"))
	requireEqual(t, shMake(-12, 0), act)
}

func Test_SH7_1(t *testing.T) {
	// number.Div divides a specified amount from the number

	act := newShopspring("12")

	act.Div(newShopspring("3"))
	requireEqual(t, shMake(4, 0), act)

	act.Div(newShopspring("-2"))
	requireEqual(t, shMake(-2, 0), act)
}

func Test_SH8_1(t *testing.T) {
	// number.Mod finds the remainder of a specified amount

	act := newShopspring("12")

	act.Mod(newShopspring("5"))
	requireEqual(t, shMake(2, 0), act)
}
