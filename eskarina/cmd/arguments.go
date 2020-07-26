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

func (args *Arguments) take() string {
	arg := (*args.list)[0]
	*args.list = (*args.list)[1:]
	return arg
}

func (args *Arguments) peek() string {
	return (*args.list)[0]
}

func (args *Arguments) count() int {
	return len(*args.list)
}

func (args *Arguments) empty() bool {
	return len(*args.list) == 0
}

func (args *Arguments) more() bool {
	return len(*args.list) > 0
}
