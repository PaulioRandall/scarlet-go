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

	type sliceable interface {
		Slice(int64, int64) value.Value
		Len() int64
	}

	if len(in) != 3 {
		setError(env, "@Slice requires three arguments")
		return
	}

	v, ok := in[0].(sliceable)
	if !ok {
		setError(env, "@Slice argument is not sliceable")
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

	type indexed interface {
		InRange(int64) bool
		At(int64) value.Value
		Len() int64
	}

	if len(in) != 2 {
		setError(env, "@At requires two arguments")
		return
	}

	v, ok := in[0].(indexed)
	if !ok {
		setError(env, "@At argument is not indexed")
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
