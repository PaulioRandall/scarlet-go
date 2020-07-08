package runtime

type Runtime struct {
	env *environment
}

func New() Runtime {
	return Runtime{}
}
