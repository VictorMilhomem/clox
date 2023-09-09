package main

import (
	"github.com/VictorMilhomem/golox/cmd/chunk"
	"github.com/VictorMilhomem/golox/cmd/vm"
)

func main() {
	ck := chunk.NewChunk()
	constant := ck.AddConstant(1.8)
	ck.WriteChunk(byte(chunk.OpConstant), 123)
	ck.WriteChunk(byte(constant), 123)

	constant = ck.AddConstant(3.2)
	ck.WriteChunk(byte(chunk.OpConstant), 123)
	ck.WriteChunk(byte(constant), 123)
	ck.WriteChunk(byte(chunk.OpAdd), 123)
	ck.WriteChunk(byte(chunk.OpReturn), 123)
	ck.DisassembleChunk("test chunk")
	vm := vm.NewVm()
	vm.Interpret(ck)
	ck.FreeChunk()
}
