package values

import "fmt"

type Value float64

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
	fmt.Printf("%g", value)
}
