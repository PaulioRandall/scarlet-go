package number

import (
	"github.com/shopspring/decimal"
)

func newShopspring(numStr string) Number {

	d, e := decimal.NewFromString(numStr)

	if e != nil {
		panic("SANITY CHECK! Unparsable number '" + numStr + "'")
	}

	return &number{d: d}
}

func newFromIntShopspring(i int64) Number {
	d := decimal.NewFromInt(i)
	return &number{d: d}
}

func newFromFloatShopspring(f float64) Number {
	d := decimal.NewFromFloat(f)
	return &number{d: d}
}

type number struct {
	d decimal.Decimal
}

func cast(n Number) *number {
	if v, ok := n.(*number); ok {
		return v
	}

	panic("SANITY CHECK! Unknown Number type")
}

func (n number) String() string {
	return n.d.String()
}

func (n number) Copy() Number {
	return &number{d: n.d}
}

func (n number) Integer() int64 {
	return n.d.IntPart()
}

func (n *number) Inc(count int64) {
	o := decimal.NewFromInt(count)
	n.d = n.d.Add(o)
}

func (n *number) Dec(count int64) {
	o := decimal.NewFromInt(count)
	n.d = n.d.Sub(o)
}

func (n *number) Neg() {
	n.d = n.d.Neg()
}

func (n *number) Add(other Number) {
	o := cast(other)
	n.d = n.d.Add(o.d)
}

func (n *number) Sub(other Number) {
	o := cast(other)
	n.d = n.d.Sub(o.d)
}

func (n *number) Mul(other Number) {
	o := cast(other)
	n.d = n.d.Mul(o.d)
}

func (n *number) Div(other Number) {
	o := cast(other)
	n.d = n.d.Div(o.d)
}

func (n *number) Mod(other Number) {
	o := cast(other)
	n.d = n.d.Mod(o.d)
}

func (n *number) Equal(other Number) bool {
	o := cast(other)
	return n.d.Equal(o.d)
}

func (n *number) NotEqual(other Number) bool {
	return !n.Equal(other)
}

func (n *number) LessThan(other Number) bool {
	o := cast(other)
	return n.d.LessThan(o.d)
}

func (n *number) LessThanOrEqual(other Number) bool {
	o := cast(other)
	return n.d.LessThanOrEqual(o.d)
}

func (n *number) MoreThan(other Number) bool {
	o := cast(other)
	return n.d.GreaterThan(o.d)
}

func (n *number) MoreThanOrEqual(other Number) bool {
	o := cast(other)
	return n.d.GreaterThanOrEqual(o.d)
}
