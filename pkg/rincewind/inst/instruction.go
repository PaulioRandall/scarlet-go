package inst

type Instruction interface {
	Code() Code
	Data() interface{}
	Begin() (int, int)
	End() (int, int)
	String() string
}
