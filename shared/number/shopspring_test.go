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

func Test_SH9_1(t *testing.T) {
	// number.Equal returns true if numbers are equal

	act := newShopspring("123.456").Equal(newShopspring("123.456"))
	require.True(t, act)

	act = newShopspring("123").Equal(newShopspring("456"))
	require.False(t, act)
}

func Test_SH10_1(t *testing.T) {
	// number.NotEqual returns true if numbers are equal

	act := newShopspring("123.456").NotEqual(newShopspring("123.456"))
	require.False(t, act)

	act = newShopspring("123").NotEqual(newShopspring("456"))
	require.True(t, act)
}

func Test_SH11_1(t *testing.T) {
	// number.Less returns true if receiver is less than the argument

	act := newShopspring("123").Less(newShopspring("456"))
	require.True(t, act)

	act = newShopspring("123").Less(newShopspring("123"))
	require.False(t, act)

	act = newShopspring("456").Less(newShopspring("123"))
	require.False(t, act)
}

func Test_SH12_1(t *testing.T) {
	// number.LessOrEqual returns true if receiver is less or equal than
	// the argument

	act := newShopspring("123").LessOrEqual(newShopspring("456"))
	require.True(t, act)

	act = newShopspring("123").LessOrEqual(newShopspring("123"))
	require.True(t, act)

	act = newShopspring("456").LessOrEqual(newShopspring("123"))
	require.False(t, act)
}

func Test_SH13_1(t *testing.T) {
	// number.More returns true if receiver is more than the argument

	act := newShopspring("123").More(newShopspring("456"))
	require.False(t, act)

	act = newShopspring("123").More(newShopspring("123"))
	require.False(t, act)

	act = newShopspring("456").More(newShopspring("123"))
	require.True(t, act)
}

func Test_SH14_1(t *testing.T) {
	// number.MoreOrEqual returns true if receiver is more or equal than
	// the argument

	act := newShopspring("123").MoreOrEqual(newShopspring("456"))
	require.False(t, act)

	act = newShopspring("123").MoreOrEqual(newShopspring("123"))
	require.True(t, act)

	act = newShopspring("456").MoreOrEqual(newShopspring("123"))
	require.True(t, act)
}
