package spells

import (
	"strconv"

	"github.com/PaulioRandall/scarlet-go/scarlet/spell"
	"github.com/PaulioRandall/scarlet-go/scarlet/value"
)

var noId = value.Ident("")

func getList(env spell.Runtime, id value.Ident) value.List {

	listVal := env.Fetch(id)
	if listVal == nil {
		setError(env, "Could not find the list '"+id.String()+"'")
		return nil
	}

	list, ok := listVal.(value.List)
	if !ok {
		setError(env, "'"+id.String()+"' is not a list")
		return nil
	}

	return list
}

func getList_Id(env spell.Runtime, idVal value.Value) (value.List, value.Ident) {

	id, ok := idVal.(value.Str)
	if !ok {
		setError(env, "Requires its first argument be an identifier")
		return nil, noId
	}

	list := getList(env, id.ToIdent())
	if list == nil {
		return nil, noId
	}

	return list, id.ToIdent()
}

func validIndex(env spell.Runtime, list value.List, idx value.Num) bool {
	if idx.Int() < 0 || idx.Int() >= int64(len(list)) {
		setError(env, "Out of range, list["+strconv.Itoa(len(list))+"],"+
			" given "+idx.String())
		return false
	}
	return true
}

func getList_Id_Idx(
	env spell.Runtime,
	in []value.Value,
) (_ value.List, _ value.Ident, _ value.Num) {

	id, ok := in[0].(value.Str)
	if !ok {
		setError(env, "Requires its first argument be an identifier")
		return
	}

	idx, ok := in[1].(value.Num)
	if !ok {
		setError(env, "Requires its second argument be an index")
		return
	}

	list := getList(env, id.ToIdent())
	if list == nil {
		return
	}

	if !validIndex(env, list, idx) {
		return
	}

	return list, id.ToIdent(), idx
}

func List_New(env spell.Runtime, in []value.Value, out *spell.Output) {
	list := make([]value.Value, len(in))
	for i, v := range in {
		list[i] = v
	}
	out.Set(0, value.List(list))
}

func List_Set(env spell.Runtime, in []value.Value, _ *spell.Output) {

	if len(in) != 3 {
		setError(env, "Three arguments required")
		return
	}

	list, id, idx := getList_Id_Idx(env, in)
	if list == nil {
		return
	}

	list[idx.Int()] = in[2]
	env.Bind(id, list)
}

func List_Get(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 2 {
		setError(env, "Two arguments required")
		return
	}

	list, _, idx := getList_Id_Idx(env, in)
	if list == nil {
		return
	}

	out.Set(0, list[idx.Int()])
}

func List_Prepend(env spell.Runtime, in []value.Value, _ *spell.Output) {

	if len(in) != 2 {
		setError(env, "Two arguments required")
		return
	}

	list, id := getList_Id(env, in[0])
	if list == nil {
		return
	}

	list = append(in[1:], list...)
	env.Bind(id, list)
}

func List_Append(env spell.Runtime, in []value.Value, _ *spell.Output) {

	if len(in) != 2 {
		setError(env, "Two arguments required")
		return
	}

	list, id := getList_Id(env, in[0])
	if list == nil {
		return
	}

	list = append(list, in[1:]...)
	env.Bind(id, list)
}

func List_Pop(env spell.Runtime, in []value.Value, out *spell.Output) {

	if len(in) != 1 {
		setError(env, "Two arguments required")
		return
	}

	list, id := getList_Id(env, in[0])
	if list == nil {
		return
	}

	if len(list) == 0 {
		setError(env, "Can't pop '"+id.String()+", it's empty")
		return
	}

	out.Set(0, list[0])
	list = list[1:]
	env.Bind(id, list)
}
