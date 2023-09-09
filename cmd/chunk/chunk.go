package chunk

import (
	"fmt"

	"github.com/VictorMilhomem/golox/cmd/values"
)

type Opcode byte

const (
	OpConstant Opcode = iota
	OpAdd
	OpSub
	OpMultiply
	OpDivide
	OpNegate
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
	return len(c.constants.GetValues()) - 1
}

func (c *Chunk) DisassembleChunk(name string) error {
	fmt.Printf("== %s ==\n", name)
	offset := 0
	var err error
	for offset < len(c.code) {
		offset, err = c.DisassembleInstruction(offset)
		if err != nil {
			return fmt.Errorf("error disassembling instruction %v", err)
		}

	}
	return nil
}

func (c *Chunk) DisassembleInstruction(offset int) (int, error) {
	fmt.Printf("%04d ", offset)

	line, ok := c.lines[offset]
	if !ok {
		fmt.Printf("    | ")
	} else {
		fmt.Printf("%4d ", line)
	}

	instruction := c.code[offset]

	switch instruction {
	case byte(OpConstant):
		return constantInstruction("OP_CONSTANT", c, offset), nil
	case byte(OpAdd):
		return simpleInstruction("OP_ADD", offset), nil
	case byte(OpSub):
		return simpleInstruction("OP_SUB", offset), nil
	case byte(OpMultiply):
		return simpleInstruction("OP_MULTIPLY", offset), nil
	case byte(OpDivide):
		return simpleInstruction("OP_DIVIDE", offset), nil
	case byte(OpNegate):
		return simpleInstruction("OP_NEGATE", offset), nil
	case byte(OpReturn):
		return simpleInstruction("OP_RETURN", offset), nil
	default:
		return offset + 1, fmt.Errorf("unknown opcode %d", instruction)
	}
}

func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

func constantInstruction(name string, chunk *Chunk, offset int) int {
	constant := chunk.code[offset+1]
	fmt.Printf("%-16s %4d '", name, constant)
	values.PrintValue(chunk.constants.GetValues()[constant])
	fmt.Println("'")
	return offset + 2
}
