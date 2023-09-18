package vm

import (
	"fmt"
	"os"
	"strconv"

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
	chunk   chunk.Chunk
	ip      byte
	stack   *stack.Stack
	globals map[string]values.Value
}

func NewVm() *VM {
	return &VM{
		chunk:   chunk.Chunk{},
		ip:      0,
		stack:   stack.New(),
		globals: make(map[string]values.Value),
	}
}

func (vm *VM) InitVM() {
	vm.resetStack()
}

func (vm *VM) FreeVM() {
	vm.chunk = chunk.Chunk{}
	vm.ip = 0
	vm.resetStack()
	vm.globals = make(map[string]values.Value)
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
		case chunk.OpPrint:
			value := vm.stack.Pop().(values.Value)
			value.PrintValue()
			fmt.Println()
		case chunk.OpPop:
			vm.stack.Pop()
		case chunk.OpDefineGlobal:
			constant := vm.readConstant()
			if values.IsString(constant) {
				vm.globals[constant.AsString()] = vm.stack.Peek().(values.Value)
			}
			vm.stack.Pop()
		case chunk.OpGetGlobal:
			name := vm.readConstant().AsString()
			val, ok := vm.globals[name]
			if !ok {
				vm.runtimeError("Undefined variable '%s'", name)
				return IntepretRuntimeError
			}
			vm.stack.Push(val)
		case chunk.OpSetGlobal:
			name := vm.readConstant().AsString()
			var ok bool
			vm.globals[name], ok = vm.stack.Peek().(values.Value)
			if !ok {
				vm.globals[name] = nil
				vm.runtimeError("Undefined variable '%s'", name)
				return IntepretRuntimeError
			}

		case chunk.OpReturn:
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
	b := vm.stack.Pop().(values.Value)
	a := vm.stack.Pop().(values.Value)
	switch {
	case values.IsNumber(b) && values.IsNumber(a):
		a := a.AsNumber()
		b := b.AsNumber()
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
	case values.IsNumber(b) && values.IsString(a):
		a := a.AsString()
		b := strconv.FormatFloat(b.AsNumber(), 'f', -1, 64)
		result := a + b
		if op == "+" {
			vm.stack.Push(values.StringVal(result))
		}
	case values.IsString(b) && values.IsNumber(a):
		a := strconv.FormatFloat(a.AsNumber(), 'f', -1, 64)
		b := b.AsString()
		result := a + b
		if op == "+" {
			vm.stack.Push(values.StringVal(result))
		}
	case values.IsString(b) && values.IsString(a):
		a := a.AsString()
		b := b.AsString()
		result := a + b
		if op == "+" {
			vm.stack.Push(values.StringVal(result))
		}

	default:
		vm.runtimeError("Operands must be numbers or strings")
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
