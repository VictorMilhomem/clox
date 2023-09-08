package chunk

import (
	"fmt"

	"github.com/VictorMilhomem/golox/cmd/values"
)

type Opcode byte

const (
	OpReturn Opcode = iota
	OpConstant
)

type Chunk struct {
	code      []byte
	constants *values.ValueArray
	lines     []int
}

func NewChunk() *Chunk {
	return &Chunk{
		code:      []byte{},
		constants: values.NewValueArray(),
		lines:     []int{},
	}
}

func (c *Chunk) WriteChunk(_byte byte, line int) error {
	c.code = append(c.code, _byte)
	c.lines = append(c.lines, line)
	return nil
}

func (c *Chunk) FreeChunk() error {
	c.code = []byte{}
	c.constants.FreeValueArray()
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

	if offset > 0 && c.lines[offset] == c.lines[offset-1] {
		fmt.Printf("    | ")
	} else {
		fmt.Printf("%4d ", c.lines[offset])
	}

	instruction := c.code[offset]

	switch instruction {
	case byte(OpReturn):
		return simpleInstruction("OP_RETURN", offset), nil
	case byte(OpConstant):
		return constantInstruction("OP_CONSTANT", c, offset), nil
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
