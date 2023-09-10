package values

import "fmt"

type ValueType int

const (
	VAL_BOOL ValueType = iota
	VAL_NIL
	VAL_NUMBER
)

type Value struct {
	kind ValueType
	as   interface{}
}

func NewValue(kind ValueType, as interface{}) Value {
	return Value{
		kind: kind,
		as:   as,
	}
}

func AsBool(value Value) bool {
	if value.kind != VAL_BOOL {
		panic("Valor não é do tipo VAL_BOOL")
	}
	return value.as.(bool)
}

func AsNil(value Value) interface{} {
	if value.kind != VAL_NIL {
		panic("Valor não é do tipo VAL_BOOL")
	}
	return nil
}

// AsNumber obtém o valor numérico subjacente de um Value
func AsNumber(value Value) float64 {
	if value.kind != VAL_NUMBER {
		panic("Valor não é do tipo VAL_NUMBER")
	}
	return value.as.(float64)
}

func IsBool(value Value) bool {
	return value.kind == VAL_BOOL
}

func IsNil(value Value) bool {
	return value.kind == VAL_NIL
}

func IsNumber(value Value) bool {
	return value.kind == VAL_NUMBER
}

type ValueArray struct {
	values []Value
}

func NewValueArray() *ValueArray {
	return &ValueArray{
		values: []Value{},
	}
}

func (v *ValueArray) WriteValueArray(value Value) error {
	v.values = append(v.values, value)
	return nil
}

func (v *ValueArray) FreeValueArray() error {
	v.values = []Value{}
	return nil
}

func (v *ValueArray) GetValues() []Value {
	return v.values
}

func PrintValue(value Value) {
	switch value.kind {
	case VAL_BOOL:
		{
			switch AsBool(value) {
			case true:
				fmt.Printf("true")
			case false:
				fmt.Printf("false")
			}
		}

	case VAL_NIL:
		fmt.Printf("%v", AsNil(value))
	case VAL_NUMBER:
		fmt.Printf("%g", AsNumber(value))
	}
}

func ValuesEqual(a Value, b Value) bool {
	if a.kind != b.kind {
		return false
	}
	switch a.kind {
	case VAL_BOOL:
		return AsBool(a) == AsBool(b)
	case VAL_NIL:
		return true
	case VAL_NUMBER:
		return AsNumber(a) == AsNumber(b)
	default:
		return false
	}
}
