package spells

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

const err_code = 1 // General program error

func setError(env spell.Runtime, m string, args ...interface{}) {
	s := fmt.Sprintf(m, args...)
	env.Fail(err_code, errors.New(s))
}

func Exit(env spell.Runtime, in []value.Value, _ *spell.Output) {

	if len(in) != 1 {
		setError(env, "@Exit requires one argument")
		return
	}

	c, ok := in[0].(value.Num)
	if !ok {
		setError(env, "@Exit requires its argument be a number")
		return
	}

	env.Exit(int(c.Int()))
}

func Panic(env spell.Runtime, in []value.Value, _ *spell.Output) {

	if len(in) != 2 {
		setError(env, "@Panic requires two arguments")
		return
	}

	c, ok := in[0].(value.Num)
	if !ok {
		setError(env, "@Panic requires its first argument be a number")
		return
	}

	m, ok := in[1].(value.Str)
	if !ok {
		setError(env, "@Panic requires its second argument be a string")
		return
	}

	fmt.Println(m)
	env.Exit(int(c.Int()))
}

func Str(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 1 {
		setError(env, "@Str requires one argument")
		return
	}

	out.Set(0, value.Str(in[0].String()))
}

func Print(env spell.Runtime, in []value.Value, _ *spell.Output) {
	for _, v := range in {
		fmt.Print(v.String())
	}
}

func Println(env spell.Runtime, in []value.Value, out *spell.Output) {
	Print(env, in, out)
	fmt.Println()
}

func ParseNum(env spell.Runtime, in []value.Value, out *spell.Output) {

	const name = "@ParseNum"

	if len(in) != 1 {
		setError(env, name+" requires one argument")
		return
	}

	v, ok := in[0].(value.Str)
	if !ok {
		setError(env, name+" argument must be a string")
	}

	n, e := strconv.ParseFloat(string(v), 64)
	if e != nil {
		out.Set(1, value.Str("Unable to parse `"+string(v)+"`"))
		return
	}

	out.Set(0, value.Num(n))
}

func PrintScope(env spell.Runtime, _ []value.Value, _ *spell.Output) {
	for id, v := range env.Scope() {
		fmt.Println(id.String() + ": " + v.String())
	}
}

func NewList(env spell.Runtime, in []value.Value, out *spell.Output) {
	list := make([]value.Value, len(in))
	for i, v := range in {
		list[i] = v
	}
	out.Set(0, value.List(list))
}

func Len(env spell.Runtime, in []value.Value, out *spell.Output) {

	type lengthy interface {
		Len() int64
	}

	if len(in) != 1 {
		setError(env, "@Len requires one argument")
		return
	}

	v, ok := in[0].(lengthy)
	if !ok {
		setError(env, "@Len argument has no length property")
		return
	}

	out.Set(0, value.Num(v.Len()))
}

func Slice(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 3 {
		setError(env, "@Slice requires three arguments")
		return
	}

	v, ok := in[0].(value.OrdCon)
	if !ok {
		setError(env, "@Slice argument is not an ordered container")
		return
	}

	size := v.Len()

	start, ok := in[1].(value.Num)
	if !ok {
		setError(env, "@Slice requires its second argument be an index")
		return
	}

	if start.Int() < 0 || start.Int() >= int64(size) {
		max := strconv.FormatInt(size, 10)
		setError(env, "Out of range, list["+max+"], given "+start.String())
		return
	}

	end, ok := in[2].(value.Num)
	if !ok {
		setError(env, "@Slice requires its third argument be an index")
		return
	}

	if end.Int() > int64(size) {
		max := strconv.FormatInt(size, 10)
		setError(env, "Out of range, sliceable["+max+"], given "+end.String())
		return
	}

	if end.Int() < start.Int() {
		setError(env, "Invalid range, sliceable["+start.String()+":"+end.String()+"]")
		return
	}

	out.Set(0, v.Slice(start.Int(), end.Int()))
}

func At(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 2 {
		setError(env, "@At requires two arguments")
		return
	}

	v, ok := in[0].(value.OrdCon)
	if !ok {
		setError(env, "@At argument is not an ordered container")
		return
	}

	idx, ok := in[1].(value.Num)
	if !ok {
		setError(env, "@At requires its second argument be an index")
		return
	}

	if !v.InRange(idx.Int()) {
		size := v.Len()
		max := strconv.FormatInt(size, 10)
		setError(env, "Out of range, indexed["+max+"], given "+idx.String())
		return
	}

	out.Set(0, v.At(idx.Int()))
}

func Index(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 2 {
		setError(env, "@Index requires two arguments")
		return
	}

	haystack, ok := in[0].(value.OrdCon)
	if !ok {
		setError(env, "@Index argument is not an ordered container")
		return
	}

	needle := in[1]
	if !haystack.CanHold(needle) {
		setError(env, "@Index: that container can't hold a '"+needle.Name()+"'")
		return
	}

	i := haystack.Index(needle)
	out.Set(0, value.Num(i))
}

func InRange(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 2 {
		setError(env, "@InRange requires two arguments")
		return
	}

	v, ok := in[0].(value.OrdCon)
	if !ok {
		setError(env, "@InRange first argument is not an ordered container")
		return
	}

	idx, ok := in[1].(value.Num)
	if !ok {
		setError(env, "@InRange second argument must be an index")
		return
	}

	inRange := v.InRange(idx.Int())
	out.Set(0, value.Bool(inRange))
}

func Push(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) < 1 {
		setError(env, "@Push requires at least one argument")
		return
	}

	c, ok := in[0].(value.OrdCon)
	if !ok {
		setError(env, "@Push first argument is not an ordered contianer")
		return
	}

	for _, v := range in[1:] {
		if !c.CanHold(v) {
			setError(env, "@Push: that container can't hold a '"+v.Name()+"'")
			return
		}
	}

	out.Set(0, c.PushFront(in[1:]...))
}

func Add(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) < 1 {
		setError(env, "@Add requires at least one argument")
		return
	}

	c, ok := in[0].(value.OrdCon)
	if !ok {
		setError(env, "@Add first argument is not aan ordered contianer")
		return
	}

	for _, v := range in[1:] {
		if !c.CanHold(v) {
			setError(env, "@Add: that container can't hold a '"+v.Name()+"'")
			return
		}
	}

	out.Set(0, c.PushBack(in[1:]...))
}

func Set(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 3 {
		setError(env, "@Set: Three arguments required")
		return
	}

	c, ok := in[0].(value.MutOrdCon)
	if !ok {
		setError(env, "@Set: first argument is not a mutable ordered contianer")
		return
	}

	k := in[1]
	if !c.CanBeKey(k) {
		setError(env, "@Set: invalid container key '"+k.String()+"'")
		return
	}

	v := in[2]
	if !c.CanHold(v) {
		setError(env, "@Set: that container can't hold a '"+v.Name()+"'")
		return
	}

	out.Set(0, c.Set(k, v))
}

func Del(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 2 {
		setError(env, "@Del: Two arguments required")
		return
	}

	c, ok := in[0].(value.Con)
	if !ok {
		setError(env, "@Del: first argument is not a contianer")
		return
	}

	k := in[1]
	if !c.CanBeKey(k) {
		setError(env, "@Del: invalid container key '"+k.String()+"'")
		return
	}

	c, v := c.Delete(k)
	out.Set(0, c)
	out.Set(0, v)
}

func Pop(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 1 {
		setError(env, "@Pop: one argument required")
		return
	}

	c, ok := in[0].(value.OrdCon)
	if !ok {
		setError(env, "@Pop: first argument is not an ordered contianer")
		return
	}

	if c.Len() == 0 {
		setError(env, "@Pop: can't pop from an empty container")
		return
	}

	c, v := c.PopFront()
	out.Set(0, c)
	out.Set(1, v)
}

func Take(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 1 {
		setError(env, "@Take: one argument required")
		return
	}

	c, ok := in[0].(value.OrdCon)
	if !ok {
		setError(env, "@Take: first argument is not an ordered contianer")
		return
	}

	if c.Len() == 0 {
		setError(env, "@Take: can't take from an empty container")
		return
	}

	c, v := c.PopBack()
	out.Set(0, c)
	out.Set(1, v)
}

func Join(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 2 {
		setError(env, "@Join: two arguments required")
		return
	}

	a, ok := in[0].(value.Joinable)
	if !ok {
		setError(env, "@Join: first argument is not a joinable value")
		return
	}

	b := in[1]
	if !a.CanJoin(b) {
		setError(env, "@Join: '"+in[0].Name()+"' can't join with a '"+b.Name()+"'")
		return
	}

	out.Set(0, a.Join(b))
}
