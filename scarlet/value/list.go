package value

import (
	"strings"
)

type List []Value

func (List) Name() string              { return "list" }
func (a List) Comparable(b Value) bool { _, ok := b.(List); return ok }

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
		if i > 2 {
			sb.WriteString("...")
			break
		}
		sb.WriteString(v.String())
	}
	sb.WriteRune(']')
	return sb.String()
}

func (a List) Len() int {
	return len(a)
}
