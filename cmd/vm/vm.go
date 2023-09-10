package vm

import (
	"fmt"
	"os"

	"github.com/VictorMilhomem/golox/cmd/chunk"
	"github.com/VictorMilhomem/golox/cmd/compiler"
	"github.com/VictorMilhomem/golox/cmd/values"
	"github.com/antlr4-go/antlr/v4"
	"github.com/golang-collections/collections/stack"
)

type InterpretResult int

const STACKMAX int = 256

const (
	IntepretOk InterpretResult = iota
	InterpretCompileError
	IntepretRuntimeError
)

type VM struct {
	chunk *chunk.Chunk
	ip    byte
	stack *stack.Stack
}

func NewVm() *VM {
	return &VM{
		chunk: chunk.NewChunk(),
		ip:    0,
		stack: stack.New(),
	}
}

func (vm *VM) InitVM() error {
	vm.resetStack()
	return nil
}

func (vm *VM) FreeVM() error {
	vm.chunk = chunk.NewChunk()
	vm.ip = 0
	vm.resetStack()
	return nil
}

func (vm *VM) resetStack() {
	vm.stack = stack.New()
}

func (vm *VM) Interpret(source string) InterpretResult {
	is := antlr.NewInputStream(source)
	ck := chunk.NewChunk()
	if !compiler.Compile(is, ck) {
		ck.FreeChunk()
		return InterpretCompileError
	}

	vm.chunk = ck
	vm.ip = vm.chunk.GetCode()[0]
	result := vm.run()

	ck.FreeChunk()
	return result
}

func (vm *VM) run() InterpretResult {
	const debug bool = true
	for {
		if debug {
			vm.chunk.DisassembleInstruction(int(vm.ip))
		}
		instruction := vm.readByte()

		switch instruction {
		case byte(chunk.OpConstant):
			{
				constant := vm.readConstant()
				vm.stack.Push(constant)
			}
		case byte(chunk.OpNil):
			vm.stack.Push(values.NewValue(values.VAL_NIL, 0))
		case byte(chunk.OpTrue):
			vm.stack.Push(values.NewValue(values.VAL_BOOL, true))
		case byte(chunk.OpFalse):
			vm.stack.Push(values.NewValue(values.VAL_BOOL, false))

		case byte(chunk.OpEqual):
			{
				b := vm.stack.Pop().(values.Value)
				a := vm.stack.Pop().(values.Value)
				vm.stack.Push(values.NewValue(values.VAL_BOOL, values.ValuesEqual(a, b)))
			}
		case byte(chunk.OpGreater):
			vm.binaryOp(">")
		case byte(chunk.OpLess):
			vm.binaryOp("<")
		case byte(chunk.OpNot):
			{
				val := values.NewValue(values.VAL_BOOL, vm.isFalsey(vm.stack.Pop().(values.Value)))
				vm.stack.Push(val)
			}
		case byte(chunk.OpAdd):
			vm.binaryOp("+")
		case byte(chunk.OpSub):
			vm.binaryOp("-")
		case byte(chunk.OpMultiply):
			vm.binaryOp("*")
		case byte(chunk.OpDivide):
			vm.binaryOp("/")
		case byte(chunk.OpNegate):
			{
				if !values.IsNumber(vm.stack.Peek().(values.Value)) {
					vm.runtimeError("operand must be a number")
				}

				value := values.AsNumber(vm.stack.Pop().(values.Value))
				vm.stack.Push(values.NewValue(values.VAL_NUMBER, -value))
			}

		case byte(chunk.OpReturn):
			value, _ := vm.stack.Pop().(values.Value)
			values.PrintValue(value)
			fmt.Println()
			return IntepretOk
		}
	}
}

func (vm *VM) runtimeError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)

	instruction := int(vm.ip) - 1
	line := vm.chunk.GetLines()[instruction]
	fmt.Fprintf(os.Stderr, "[line %d] in script\n", line)

	// Reset your stack if needed
	vm.resetStack()
}

func (vm *VM) binaryOp(op string) {
	if !values.IsNumber(vm.stack.Peek().(values.Value)) {
		vm.runtimeError("operands must be numbers")
	} else {
		b := values.AsNumber(vm.stack.Pop().(values.Value))
		if values.IsNumber(vm.stack.Peek().(values.Value)) {
			a := values.AsNumber(vm.stack.Pop().(values.Value))
			switch op {
			case "+":
				vm.stack.Push(values.NewValue(values.VAL_NUMBER, a+b))
			case "-":
				vm.stack.Push(values.NewValue(values.VAL_NUMBER, a-b))
			case "*":
				vm.stack.Push(values.NewValue(values.VAL_NUMBER, a*b))
			case "/":
				vm.stack.Push(values.NewValue(values.VAL_NUMBER, a/b))
			case ">":
				vm.stack.Push(values.NewValue(values.VAL_BOOL, a > b))
			case "<":
				vm.stack.Push(values.NewValue(values.VAL_BOOL, a < b))
			}
		}
	}
}

func (vm *VM) isFalsey(value values.Value) bool {
	return values.IsNil(value) || (values.IsBool(value) && !values.AsBool(value))
}

func (vm *VM) readByte() byte {
	if vm.ip >= byte(len(vm.chunk.GetCode())) {
		return 0
	}

	byteToRead := vm.chunk.GetCode()[vm.ip]
	vm.ip++
	return byteToRead
}

func (vm *VM) readConstant() values.Value {
	return vm.chunk.GetConstants().GetValues()[vm.readByte()]
}
