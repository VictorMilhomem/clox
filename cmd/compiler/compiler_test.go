package compiler

import (
	"bytes"
	"testing"

	"github.com/VictorMilhomem/golox/cmd/chunk"
)

func TestCompile(t *testing.T) {
	source := "1 + 1.2\x00"
	ck := chunk.NewChunk()

	success := true
	if !Compile(source, ck) {
		ck.FreeChunk()
		success = false
	}

	// Verifique se a compilação foi bem-sucedida
	if !success {
		t.Error("Compile failed")
	}

	// Insira as asserções adicionais aqui para verificar se o chunk contém o código esperado.
	// Por exemplo, você pode usar a função chunk.Disassemble para inspecionar o chunk gerado.

	// Exemplo de asserção (substitua pelo seu próprio teste):
	expectedInstructions := []byte{byte(chunk.OpConstant), byte(chunk.OpAdd), byte(chunk.OpConstant), byte(chunk.OpReturn)} // Substitua pelo seu código esperado
	if !bytes.Equal(ck.GetCode(), expectedInstructions) {
		t.Errorf("Expected chunk code: %v, got: %v", expectedInstructions, ck.GetCode())
	}
}
