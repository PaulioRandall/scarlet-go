package spells

import (
	"strconv"

	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

var zero = value.Num(0)
var noId = value.Ident("")

func getListId_Idx(env spell.Runtime, in []value.Value, errPre string) (value.Ident, value.Num) {
	id, ok := in[0].(value.Str)
	if !ok {
		setError(env, errPre+": requires its first argument be an identifier")
		return noId, zero
	}

	idx, ok := in[1].(value.Num)
	if !ok {
		setError(env, errPre+": requires its second argument be an index")
		return noId, zero
	}

	return id.ToIdent(), idx
}

func getList(env spell.Runtime, id value.Ident, errPre string) value.List {

	listVal := env.Fetch(id)
	if listVal == nil {
		setError(env, errPre+": Could not find the list '"+id.String()+"'")
		return nil
	}

	list, ok := listVal.(value.List)
	if !ok {
		setError(env, errPre+": '"+id.String()+"' is not a list")
		return nil
	}

	return list
}

func validIndex(env spell.Runtime, list value.List, idx value.Num, errPre string) bool {
	if idx.Int() < 0 || idx.Int() >= int64(len(list)) {
		setError(env, errPre+": Out of range, list["+
			strconv.Itoa(len(list))+"], given "+idx.String())
		return false
	}
	return true
}

func List_New(env spell.Runtime, in []value.Value, out *spell.Output) {
	list := make([]value.Value, len(in))
	for i, v := range in {
		list[i] = v
	}
	out.Set(0, value.List(list))
}

func List_Set(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 3 {
		setError(env, "@list.Set requires three arguments")
		return
	}

	listId, idx := getListId_Idx(env, in, "@list.Set")

	list := getList(env, listId, "@list.Set")
	if list == nil {
		return
	}

	if !validIndex(env, list, idx, "@list.Set") {
		return
	}

	list[idx.Int()] = in[2]
	env.Bind(listId, list)
}

func List_Get(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 2 {
		setError(env, "@list.Get requires two arguments")
		return
	}

	listId, idx := getListId_Idx(env, in, "@list.Get")

	list := getList(env, listId, "@list.Get")
	if list == nil {
		return
	}

	if !validIndex(env, list, idx, "@list.Get") {
		return
	}

	out.Set(0, list[idx.Int()])
}
