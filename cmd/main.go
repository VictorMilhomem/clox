package main

import (
	"github.com/VictorMilhomem/golox/cmd/chunk"
)

func main() {
	ck := chunk.NewChunk()
	constant := ck.AddConstant(1.2)
	ck.WriteChunk(byte(chunk.OpConstant))
	ck.WriteChunk(byte(constant))
	ck.WriteChunk(byte(chunk.OpReturn))
	ck.DisassembleChunk("test chunk")
	ck.FreeChunk()
}
