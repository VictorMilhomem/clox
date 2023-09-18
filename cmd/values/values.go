package values

import (
	"fmt"
	"strings"
)

type (
	NilVal    struct{}
	BoolVal   bool
	NumberVal float64
	ValueType int
	StringVal string
)

const (
	VAL_BOOL ValueType = iota
	VAL_NIL
	VAL_NUMBER
	VAL_OBJ
	VAL_STR
)

type Value interface {
	Type() ValueType
	AsBoolean() bool
	AsNumber() float64
	AsString() string
	PrintValue()
}

func IsNumber(v Value) bool {
	return v.Type() == VAL_NUMBER
}

func IsBool(v Value) bool {
	return v.Type() == VAL_BOOL
}

func IsNil(v Value) bool {
	return v.Type() == VAL_NIL
}

func IsObj(v Value) bool {
	return v.Type() == VAL_OBJ
}

func IsString(v Value) bool {
	return v.Type() == VAL_STR
}

/*================NIL VAL======================*/
func (nv NilVal) Type() ValueType {
	return VAL_NIL
}

func (nv NilVal) AsBoolean() bool {
	panic("nil value is not a boolean!")
}

func (nv NilVal) AsNumber() float64 {
	panic("nil value is not a number!")
}

func (nv NilVal) AsString() string {
	panic("number value is not a boolean!")
}

func (nv NilVal) PrintValue() {
	fmt.Print("nil")
}

/*================BOOL VAL======================*/

func (bv BoolVal) Type() ValueType {
	return VAL_BOOL
}

func (bv BoolVal) AsBoolean() bool {
	return bool(bv)
}

func (bv BoolVal) AsNumber() float64 {
	panic("bool value is not a number!")
}

func (bv BoolVal) AsString() string {
	panic("bool value is not a obj!")
}

func (bv BoolVal) PrintValue() {
	fmt.Printf("%t", bool(bv))
}

/*================NUMBER VAL======================*/

func (nv NumberVal) Type() ValueType {
	return VAL_NUMBER
}

func (nv NumberVal) AsBoolean() bool {
	panic("number value is not a boolean!")
}

func (nv NumberVal) AsNumber() float64 {
	return float64(nv)
}

func (nv NumberVal) AsString() string {
	panic("number value is not a boolean!")
}

func (nv NumberVal) PrintValue() {
	fmt.Printf("%g", float64(nv))
}

/*================OBJ VAL======================*/
func (str StringVal) Type() ValueType {
	return VAL_STR
}

func (str StringVal) AsBoolean() bool {
	panic("object value is not a boolean!")
}

func (str StringVal) AsNumber() float64 {
	panic("object value is not a number!")
}

func (str StringVal) AsString() string {
	return string(str)
}

func (str StringVal) PrintValue() {
	fmt.Printf("%v", str)
}

/*================VALUE ARRAY======================*/

type ValueArray struct {
	Values []Value
}

func NewValueArray() *ValueArray {
	return &ValueArray{}
}

func (a ValueArray) Count() int {
	return len(a.Values)
}

func (a *ValueArray) WriteValueArray(v Value) error {
	a.Values = append(a.Values, v)
	return nil
}

func (a *ValueArray) FreeValueArray() {
	a.Values = []Value{}
}

func ValuesEqual(a Value, b Value) bool {
	if a.Type() != b.Type() {
		return false
	}
	switch a.Type() {
	case VAL_BOOL:
		return a.AsBoolean() == b.AsBoolean()
	case VAL_NIL:
		return true
	case VAL_NUMBER:
		return a.AsNumber() == b.AsNumber()
	case VAL_OBJ:
		vala := a.AsString()
		valb := b.AsString()
		comp := strings.Compare(vala, valb)
		if comp == 0 {
			return true
		}
		return false

	default:
		return false
	}
}
