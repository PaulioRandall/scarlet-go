package value

import (
	"strings"
)

type List []Value

func (List) Name() string              { return "list" }
func (a List) Comparable(b Value) bool { _, ok := b.(List); return ok }

func (a List) Len() int64                       { return int64(len(a)) }
func (a List) Slice(start, end int64) Container { return a[start:end] }
func (a List) CanHold(v Value) bool             { return true }
func (a List) InRange(idx int64) bool {
	return idx >= 0 && idx < int64(len(a))
}
func (a List) At(idx int64) Value           { return a[idx] }
func (a List) Prepend(v ...Value) Container { return append(List(v), a...) }
func (a List) Append(v ...Value) Container  { return append(a, List(v)...) }

func (a List) CanBeKey(v Value) bool {
	i, ok := v.(Num)
	return ok && a.InRange(i.Int())
}
func (a List) Set(i Value, v Value) MutContainer {
	a[i.(Num).Int()] = v
	return a
}

func (a List) Equal(b Value) bool {
	if !a.Comparable(b) {
		return false
	}

	o := b.(List)
	if len(a) != len(o) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !a[i].Equal(o[i]) {
			return false
		}
	}

	return true
}

func (a List) String() string {
	sb := strings.Builder{}
	sb.WriteRune('[')
	for i, v := range a {
		if i > 0 {
			sb.WriteString(", ")
		}
		if i > 4 {
			sb.WriteString("...")
			break
		}
		sb.WriteString(v.String())
	}
	sb.WriteRune(']')
	return sb.String()
}
