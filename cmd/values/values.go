package values

import "fmt"

type ValueType int

const (
	VAL_BOOL ValueType = iota
	VAL_NIL
	VAL_NUMBER
)

type Value interface { // N.B. go doesn't have sum-types, so we use an interface; this is more verbose
	Type() ValueType
	AsBoolean() bool
	AsNumber() float64
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

type NilVal struct{}

func (nv NilVal) Type() ValueType {
	return VAL_NIL
}

func (nv NilVal) AsBoolean() bool {
	panic("nil value is not a boolean!")
}

func (nv NilVal) AsNumber() float64 {
	panic("nil value is not a number!")
}

func (nv NilVal) PrintValue() {
	fmt.Print("nil")
}

type BoolVal bool

func (bv BoolVal) Type() ValueType {
	return VAL_BOOL
}

func (bv BoolVal) AsBoolean() bool {
	return bool(bv)
}

func (bv BoolVal) AsNumber() float64 {
	panic("bool value is not a number!")
}

func (bv BoolVal) PrintValue() {
	fmt.Printf("%t", bool(bv))
}

type NumberVal float64

func (nv NumberVal) Type() ValueType {
	return VAL_NUMBER
}

func (nv NumberVal) AsBoolean() bool {
	panic("number value is not a boolean!")
}

func (nv NumberVal) AsNumber() float64 {
	return float64(nv)
}

func (nv NumberVal) PrintValue() {
	fmt.Printf("%g", float64(nv))
}

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
	default:
		return false
	}
}
