package compiler

import (
	"fmt"
	"math"
	"strconv"

	"github.com/VictorMilhomem/golox/cmd/chunk"
	"github.com/VictorMilhomem/golox/cmd/values"
)

const UINT8_MAX = math.MaxUint8

type Compiler struct {
	parser         *Parser
	compilingChunk *chunk.Chunk
}

func NewCompiler(parser *Parser, scanner *Scanner) *Compiler {
	return &Compiler{
		parser: parser,
	}
}

type Parser struct {
	current   Token
	previous  Token
	hadError  bool
	panicMode bool
	scanner   *Scanner
}

func NewParser(scanner *Scanner) *Parser {
	return &Parser{
		scanner: scanner,
	}
}

func (p *Parser) advance() {
	p.previous = p.current

	for {
		p.current = p.scanner.scanToken()
		if p.current.kind != TOKEN_ERROR {
			break
		}
		p.errorAtCurrent(p.current.start)
	}
}

func (p *Parser) errorAtCurrent(message string) {
	p.errorAt(p.current, message)
}

func (p *Parser) err(message string) {
	p.errorAt(p.previous, message)
}

func (p *Parser) errorAt(token Token, message string) {
	if p.panicMode {
		return
	}
	p.panicMode = true
	fmt.Printf("[line %d] error", token.line)

	if token.kind == TOKEN_EOF {
		fmt.Printf("[line %d] error", token.line)
	} else if token.kind == TOKEN_ERROR {
		// nothing
	} else {
		fmt.Printf(" at '%s'", string(token.start[:token.length]))
	}

	fmt.Printf(": %s\n", message)
	p.hadError = true
}

func (p *Parser) consume(kind TokenType, message string) {
	if p.current.kind == kind {
		p.advance()
		return
	}

	p.errorAtCurrent(message)
}

func (c *Compiler) emitByte(_byte byte) {
	c.compilingChunk.WriteChunk(_byte, c.parser.previous.line)
}

func (c *Compiler) emitBytes(byte1, byte2 byte) {
	c.emitByte(byte1)
	c.emitByte(byte2)
}

func (c *Compiler) emitReturn() {
	c.emitByte(byte(chunk.OpReturn))
}

func (c *Compiler) emitConstant(value values.Value) {
	c.emitBytes(byte(chunk.OpConstant), c.makeConstant(value))
}

func (c *Compiler) endCompiler() {
	c.emitReturn()
}

func (c *Compiler) makeConstant(value values.Value) byte {
	constant := c.compilingChunk.AddConstant(value)
	if constant > UINT8_MAX {
		c.parser.err("too many constants in one chunk")
		return 0
	}
	return byte(constant)
}

func (c *Compiler) number() {
	value, err := strconv.ParseFloat(c.parser.previous.start, 64)
	if err != nil {
		fmt.Errorf("error parsing float")
		return
	}
	c.emitConstant(values.Value(value))
}

func Compile(source string, chunk *chunk.Chunk) bool {
	scanner := NewScanner([]rune(source))
	parser := NewParser(scanner)
	compiler := NewCompiler(parser, scanner)
	compiler.compilingChunk = chunk
	parser.hadError = false
	parser.panicMode = false
	parser.advance()
	parser.consume(TOKEN_EOF, "Expect end of expression.")
	compiler.endCompiler()
	return !parser.hadError
}
