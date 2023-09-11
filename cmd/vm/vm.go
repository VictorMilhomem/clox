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

const (
	IntepretOk InterpretResult = iota
	InterpretCompileError
	IntepretRuntimeError
)

type VM struct {
	chunk chunk.Chunk
	ip    byte
	stack *stack.Stack
}

func NewVm() *VM {
	return &VM{
		chunk: chunk.Chunk{},
		ip:    0,
		stack: stack.New(),
	}
}

func (vm *VM) InitVM() {
	vm.resetStack()
}

func (vm *VM) FreeVM() {
	vm.chunk = chunk.Chunk{}
	vm.ip = 0
	vm.resetStack()
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

	vm.chunk = *ck
	vm.ip = 0
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

		switch chunk.Opcode(instruction) {
		case chunk.OpConstant:
			{
				constant := vm.readConstant()
				vm.stack.Push(constant)
			}
		case chunk.OpNil:
			vm.stack.Push(values.NilVal(struct{}{}))
		case chunk.OpTrue:
			vm.stack.Push(values.BoolVal(true))
		case chunk.OpFalse:
			vm.stack.Push(values.BoolVal(false))

		case chunk.OpEqual:
			{
				b := vm.stack.Pop().(values.Value)
				a := vm.stack.Pop().(values.Value)
				vm.stack.Push(values.BoolVal(values.ValuesEqual(a, b)))
			}
		case chunk.OpGreater:
			vm.binaryOp(">")
		case chunk.OpLess:
			vm.binaryOp("<")
		case chunk.OpNot:
			{
				val := values.BoolVal(vm.isFalsey(vm.stack.Pop().(values.Value)))
				vm.stack.Push(val)
			}
		case chunk.OpAdd:
			vm.binaryOp("+")
		case chunk.OpSub:
			vm.binaryOp("-")
		case chunk.OpMultiply:
			vm.binaryOp("*")
		case chunk.OpDivide:
			vm.binaryOp("/")
		case chunk.OpNegate:
			{
				if !values.IsNumber(vm.stack.Peek().(values.Value)) {
					vm.runtimeError("operand must be a number")
				}

				value := vm.stack.Pop().(values.Value).AsNumber()
				vm.stack.Push(values.NumberVal(-value))
			}

		case chunk.OpReturn:
			value, _ := vm.stack.Pop().(values.Value)
			value.PrintValue()
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
		b := vm.stack.Pop().(values.Value).AsNumber()
		if values.IsNumber(vm.stack.Peek().(values.Value)) {
			a := vm.stack.Pop().(values.Value).AsNumber()
			switch op {
			case "+":
				vm.stack.Push(values.NumberVal(a + b))
			case "-":
				vm.stack.Push(values.NumberVal(a - b))
			case "*":
				vm.stack.Push(values.NumberVal(a * b))
			case "/":
				vm.stack.Push(values.NumberVal(a / b))
			case ">":
				vm.stack.Push(values.BoolVal(a > b))
			case "<":
				vm.stack.Push(values.BoolVal(a < b))
			}
		}
	}
}

func (vm *VM) isFalsey(value values.Value) bool {
	return values.IsNil(value) || (values.IsBool(value) && !value.AsBoolean())
}

func (vm *VM) readByte() byte {
	byteToRead := vm.chunk.GetCode()[vm.ip]
	vm.ip++
	return byteToRead
}

func (vm *VM) readConstant() values.Value {
	return vm.chunk.GetConstants().Values[vm.readByte()]
}
