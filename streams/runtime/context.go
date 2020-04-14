package runtime

// context represents the current executing context. It contains all state
// available to the current scope such as available variables.
type context struct {
	vars map[string]Value
}

// String returns a human readable string representation of the context.
func (ctx *context) String() (s string) {

	const NEWLINE = "\n"
	const TAB = "\t"

	s += "variable:" + NEWLINE

	if len(ctx.vars) == 0 {
		s += TAB + "(Empty)" + NEWLINE
		return
	}

	for id, v := range ctx.vars {
		s += TAB + id + v.String() + NEWLINE
	}

	return
}

// get returns the value assigned to a specified variable. If the ID does not
// exist an empty value is returned.
func (ctx *context) get(id string) (_ Value) {

	if v, ok := ctx.vars[id]; ok {
		return v
	}

	return
}

// expect returns the value assigned to a specified variable. If the ID does
// not exist a panic ensues.
func (ctx *context) expect(id string) Value {
	v := ctx.get(id)

	if v == nil {
		//panic(newErr("Expected variable '%v'", id))
		panic(string("Expected variable '" + id + "'"))
	}

	return v
}

// set creates or updates a variable. Passing a VOID value deletes the entry if
// it exists.
func (ctx *context) set(id string, v Value) {

	if _, ok := v.(Void); ok {
		delete(ctx.vars, id)
		return
	}

	ctx.vars[id] = v
}
