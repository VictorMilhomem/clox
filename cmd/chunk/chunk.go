package chunk

import (
	"fmt"

	"github.com/VictorMilhomem/golox/cmd/values"
)

type Opcode byte

const (
	OpConstant Opcode = iota
	OpNil
	OpTrue
	OpFalse
	OpEqual
	OpGreater
	OpLess
	OpAdd
	OpSub
	OpMultiply
	OpDivide
	OpNot
	OpNegate
	OpPrint
	OpPop
	OpDefineGlobal
	OpGetGlobal
	OpSetGlobal
	OpReturn
)

type Chunk struct {
	code      []byte
	constants *values.ValueArray
	lines     map[int]int
}

func NewChunk() *Chunk {
	return &Chunk{
		code:      []byte{},
		constants: values.NewValueArray(),
		lines:     make(map[int]int),
	}
}

func (c *Chunk) GetCode() []byte {
	return c.code
}

func (c *Chunk) GetConstants() *values.ValueArray {
	return c.constants
}

func (c *Chunk) GetLines() map[int]int {
	return c.lines
}

func (c *Chunk) WriteChunk(_byte byte, line int) error {
	c.code = append(c.code, _byte)
	c.lines[len(c.code)-1] = line
	return nil
}

func (c *Chunk) FreeChunk() error {
	c.code = []byte{}
	c.constants.FreeValueArray()
	c.lines = make(map[int]int)
	return nil
}

func (c *Chunk) AddConstant(value values.Value) int {
	c.constants.WriteValueArray(value)
	return len(c.constants.Values) - 1
}

func (c *Chunk) DisassembleChunk(name string) error {
	fmt.Printf("== %s ==\n", name)
	offset := 0
	for offset < len(c.code) {
		offset = c.DisassembleInstruction(offset)
	}
	return nil
}

func (c *Chunk) DisassembleInstruction(offset int) int {
	fmt.Printf("%04d ", offset)

	if offset > 0 && c.lines[offset] == c.lines[offset-1] {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", c.lines[offset])
	}

	instruction := c.code[offset]

	switch instruction {
	case byte(OpConstant):
		return constantInstruction("OP_CONSTANT", c, offset)
	case byte(OpNil):
		return simpleInstruction("OP_NIL", offset)
	case byte(OpTrue):
		return simpleInstruction("OP_TRUE", offset)
	case byte(OpFalse):
		return simpleInstruction("OP_FALSE", offset)
	case byte(OpEqual):
		return simpleInstruction("OP_EQUAL", offset)
	case byte(OpGreater):
		return simpleInstruction("OP_GREATER", offset)
	case byte(OpLess):
		return simpleInstruction("OP_LESS", offset)
	case byte(OpAdd):
		return simpleInstruction("OP_ADD", offset)
	case byte(OpSub):
		return simpleInstruction("OP_SUB", offset)
	case byte(OpMultiply):
		return simpleInstruction("OP_MULTIPLY", offset)
	case byte(OpDivide):
		return simpleInstruction("OP_DIVIDE", offset)
	case byte(OpNegate):
		return simpleInstruction("OP_NEGATE", offset)
	case byte(OpNot):
		return simpleInstruction("OP_NOT", offset)
	case byte(OpPrint):
		return simpleInstruction("OP_PRINT", offset)
	case byte(OpPop):
		return simpleInstruction("OP_POP", offset)
	case byte(OpDefineGlobal):
		return constantInstruction("OP_DEFINE_GLOBAL", c, offset)
	case byte(OpGetGlobal):
		return constantInstruction("OP_GET_GLOBAL", c, offset)
	case byte(OpSetGlobal):
		return constantInstruction("OP_SET_GLOBAL", c, offset)
	case byte(OpReturn):
		return simpleInstruction("OP_RETURN", offset)
	default:
		return offset + 1
	}
}

func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

func constantInstruction(name string, chunk *Chunk, offset int) int {
	constant := chunk.code[offset+1]
	fmt.Printf("%-16s %4d '", name, constant)
	chunk.constants.Values[constant].PrintValue()
	fmt.Println("'")
	return offset + 2
}
