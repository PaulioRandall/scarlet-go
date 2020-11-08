package value

type Str string

func (Str) Name() string              { return "string" }
func (a Str) Comparable(b Value) bool { _, ok := b.(Str); return ok }
func (a Str) Equal(b Value) bool {
	return a.Comparable(b) && a == b.(Str)
}

func (a Str) String() string { return string(a) }
func (a Str) ToIdent() Ident { return Ident(string(a)) }

func (a Str) Len() int64                   { return int64(len(string(a))) }
func (a Str) Slice(start, end int64) Value { return a[start:end] }

func (a Str) CanHold(v Value) bool   { _, ok := v.(Str); return ok }
func (a Str) InRange(idx int64) bool { return idx >= 0 && idx < int64(len(a)) }
func (a Str) At(idx int64) Value     { return Str(string([]rune(string(a))[idx])) }
func (a Str) Prepend(v Value) Value  { return v.(Str) + a }
func (a Str) Append(v Value) Value   { return a + v.(Str) }
