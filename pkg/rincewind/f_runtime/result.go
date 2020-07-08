package runtime

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/number"
	"github.com/PaulioRandall/scarlet-go/pkg/rincewind/perror"
)

type resultType int

const (
	RT_UNDEFINED resultType = iota
	RT_BOOL
	RT_NUMBER
	RT_STRING
)

var resultTypes = map[resultType]string{
	RT_BOOL:   "bool",
	RT_NUMBER: "number",
	RT_STRING: "string",
}

func (rt resultType) String() string {
	return resultTypes[rt]
}

type result struct {
	ty  resultType
	val interface{}
}

func (r result) is(ty resultType) bool {
	return r.ty == ty
}

func (r result) Bool() (bool, bool) {

	if r.ty == RT_BOOL {
		return r.val.(bool), true
	}

	return false, false
}

func (r result) Num() (number.Number, bool) {

	if r.ty == RT_NUMBER {
		return r.val.(number.Number), true
	}

	return nil, false
}

func (r result) Str() (string, bool) {

	if r.ty == RT_STRING {
		return r.val.(string), true
	}

	return "", false
}

func (r result) String() string {
	return fmt.Sprintf("%v", r.val)
}

func resultTypeOf(v interface{}) resultType {

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
