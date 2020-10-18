package cmd

import (
	"strings"
)

type Args struct {
	list *[]string
}

func NewArgs(args []string) Args {
	if args == nil {
		panic("Nil args not allowed")
	}
	return Args{
		list: &args,
	}
}

func (args *Args) shift() string {
	arg := (*args.list)[0]
	*args.list = (*args.list)[1:]
	return arg
}

func (args *Args) shiftDefault(def string) string {

	if args.empty() {
		return def
	}

	arg := (*args.list)[0]
	*args.list = (*args.list)[1:]
	return arg
}

func (args *Args) accept(s string) bool {

	if args.peek() == s {
		args.shift()
		return true
	}

	return false
}

func (args *Args) peek() string {
	return (*args.list)[0]
}

func (args *Args) count() int {
	return len(*args.list)
}

func (args *Args) empty() bool {
	return len(*args.list) == 0
}

func (args *Args) more() bool {
	return len(*args.list) > 0
}

func (args *Args) isOption() bool {
	return strings.HasPrefix(args.peek(), "-")
}
