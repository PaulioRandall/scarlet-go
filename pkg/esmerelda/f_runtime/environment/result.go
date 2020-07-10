package environment

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/number"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/perror"
)

type ResultType int

const (
	RT_UNDEFINED ResultType = iota
	RT_BOOL
	RT_NUMBER
	RT_STRING
)

var resultTypes = map[ResultType]string{
	RT_BOOL:   "bool",
	RT_NUMBER: "number",
	RT_STRING: "string",
}

func (rt ResultType) String() string {
	return resultTypes[rt]
}

type Result struct {
	ty  ResultType
	val interface{}
}

func (r Result) is(ty ResultType) bool {
	return r.ty == ty
}

func (r Result) Bool() (bool, bool) {

	if r.ty == RT_BOOL {
		return r.val.(bool), true
	}

	return false, false
}

func (r Result) Num() (number.Number, bool) {

	if r.ty == RT_NUMBER {
		return r.val.(number.Number), true
	}

	return nil, false
}

func (r Result) Str() (string, bool) {

	if r.ty == RT_STRING {
		return r.val.(string), true
	}

	return "", false
}

func (r Result) String() string {
	return fmt.Sprintf("%v", r.val)
}

func resultTypeOf(v interface{}) ResultType {

	switch v.(type) {
	case bool:
		return RT_BOOL

	case number.Number:
		return RT_NUMBER

	case string:
		return RT_STRING
	}

	perror.Panic("No result type for %v", v)
	return RT_UNDEFINED
}
