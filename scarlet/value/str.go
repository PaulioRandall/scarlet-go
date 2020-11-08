package value

import (
	"strings"
)

type Str string

func (Str) Name() string              { return "string" }
func (a Str) Comparable(b Value) bool { _, ok := b.(Str); return ok }
func (a Str) Equal(b Value) bool {
	return a.Comparable(b) && a == b.(Str)
}

func (a Str) String() string { return string(a) }
func (a Str) ToIdent() Ident { return Ident(string(a)) }

func (a Str) Len() int64                       { return int64(len(string(a))) }
func (a Str) Slice(start, end int64) Container { return a[start:end] }

func (a Str) CanHold(v Value) bool   { _, ok := v.(Str); return ok }
func (a Str) InRange(idx int64) bool { return idx >= 0 && idx < int64(len(a)) }
func (a Str) At(idx int64) Value     { return Str(string([]rune(string(a))[idx])) }

func (a Str) Prepend(v ...Value) Container {
	sb := strings.Builder{}
	for _, s := range v {
		sb.WriteString(string(s.(Str)))
	}
	sb.WriteString(string(a))
	return Str(sb.String())
}

func (a Str) Append(v ...Value) Container {
	sb := strings.Builder{}
	sb.WriteString(string(a))
	for _, s := range v {
		sb.WriteString(string(s.(Str)))
	}
	return Str(sb.String())
}
