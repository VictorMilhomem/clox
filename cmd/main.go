package main

import (
	"github.com/VictorMilhomem/golox/cmd/chunk"
)

func main() {
	ck := chunk.NewChunk()
	ck.WriteChunk(byte(chunk.OpReturn))
	ck.DisassembleChunk("test chunk")
	ck.FreeChunk()
}
