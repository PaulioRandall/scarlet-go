package cmd

type Arguments struct {
	list *[]string
}

func NewArgs(args []string) Arguments {

	if args == nil {
		panic("Nil args not allowed")
	}

	return Arguments{
		list: &args,
	}
}

func (args *Arguments) Shift() string {
	arg := (*args.list)[0]
	*args.list = (*args.list)[1:]
	return arg
}

func (args *Arguments) Peek() string {
	return (*args.list)[0]
}

func (args *Arguments) Count() int {
	return len(*args.list)
}

func (args *Arguments) Empty() bool {
	return len(*args.list) == 0
}

func (args *Arguments) More() bool {
	return len(*args.list) > 0
}
