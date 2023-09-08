package chunk

import (
	"fmt"
)

type Opcode byte

const (
	OpReturn Opcode = iota
)

type Chunk struct {
	code []byte
}

func NewChunk() *Chunk {
	return &Chunk{
		code: []byte{},
	}
}

func (c *Chunk) WriteChunk(_byte byte) error {
	c.code = append(c.code, _byte)
	return nil
}

func (c *Chunk) FreeChunk() error {
	c.code = []byte{}
	return nil
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
	instruction := c.code[offset]

	switch instruction {
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
