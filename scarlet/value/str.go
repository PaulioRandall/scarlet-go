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

func (a Str) Len() int64 {
	return int64(len([]rune(string(a))))
}
func (a Str) Slice(start, end int64) OrdCon { return a[start:end] }

func (a Str) CanBeKey(v Value) bool {
	i, ok := v.(Num)
	return ok && a.InRange(i.Int())
}
func (a Str) CanHold(v Value) bool   { _, ok := v.(Str); return ok }
func (a Str) InRange(idx int64) bool { return idx >= 0 && idx < a.Len() }
func (a Str) At(idx int64) Value     { return Str(string([]rune(string(a))[idx])) }

func (a Str) PushFront(v ...Value) OrdCon {
	sb := strings.Builder{}
	for _, s := range v {
		sb.WriteString(string(s.(Str)))
	}
	sb.WriteString(string(a))
	return Str(sb.String())
}

func (a Str) PushBack(v ...Value) OrdCon {
	sb := strings.Builder{}
	sb.WriteString(string(a))
	for _, s := range v {
		sb.WriteString(string(s.(Str)))
	}
	return Str(sb.String())
}

func (a Str) Delete(v Value) (Con, Value) {
	i := v.(Num).Int()
	s := []rune(string(a))
	r := s[i]
	s = append(s[:i], s[i+1:]...)
	return Str(string(s)), Str(string(r))
}

func (a Str) PopFront() (OrdCon, Value) {
	v := []rune(string(a))
	return Str(string(v[1:])), Str(string(v[0]))
}

func (a Str) PopBack() (OrdCon, Value) {
	last := a.Len() - 1
	v := []rune(string(a))
	return Str(string(v[:last])), Str(string(v[last]))
}
