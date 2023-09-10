package compiler

import (
	"fmt"
	"math"

	"github.com/antlr4-go/antlr/v4"

	"github.com/VictorMilhomem/golox/cmd/chunk"
	"github.com/VictorMilhomem/golox/cmd/values"
)

const UINT8_MAX = math.MaxUint8

type Parser struct {
	current   antlr.Token
	previous  antlr.Token
	hadError  bool
	panicMode bool
	scanner   *Lexer
}

func NewParser() *Parser {
	return &Parser{}
}

type Precedence int

const (
	PREC_NONE       = iota
	PREC_ASSIGNMENT // =
	PREC_OR         // or
	PREC_AND        // and
	PREC_EQUALITY   // == !=
	PREC_COMPARISON // < > <= >=
	PREC_TERM       // + -
	PREC_FACTOR     // * /
	PREC_UNARY      // ! -
	PREC_CALL       // . ()
	PREC_PRIMARY
)

type ParseFn func()

type ParseRule struct {
	prefix     ParseFn
	infix      ParseFn
	precedence Precedence
}

func getRule(kind int) ParseRule {
	rules := map[int]ParseRule{
		TOKEN_LEFT_PAREN:    {grouping, nil, PREC_NONE},
		TOKEN_RIGHT_PAREN:   {nil, nil, PREC_NONE},
		TOKEN_LEFT_BRACE:    {nil, nil, PREC_NONE},
		TOKEN_RIGHT_BRACE:   {nil, nil, PREC_NONE},
		TOKEN_COMMA:         {nil, nil, PREC_NONE},
		TOKEN_DOT:           {nil, nil, PREC_NONE},
		TOKEN_MINUS:         {unary, binary, PREC_TERM},
		TOKEN_PLUS:          {nil, binary, PREC_TERM},
		TOKEN_SEMICOLON:     {nil, nil, PREC_NONE},
		TOKEN_SLASH:         {nil, binary, PREC_FACTOR},
		TOKEN_STAR:          {nil, binary, PREC_FACTOR},
		TOKEN_BANG:          {unary, nil, PREC_NONE},
		TOKEN_BANG_EQUAL:    {nil, binary, PREC_NONE},
		TOKEN_EQUAL:         {nil, nil, PREC_NONE},
		TOKEN_EQUAL_EQUAL:   {nil, binary, PREC_NONE},
		TOKEN_GREATER:       {nil, binary, PREC_NONE},
		TOKEN_GREATER_EQUAL: {nil, binary, PREC_NONE},
		TOKEN_LESS:          {nil, binary, PREC_NONE},
		TOKEN_LESS_EQUAL:    {nil, binary, PREC_NONE},
		TOKEN_IDENTIFIER:    {nil, nil, PREC_NONE},
		TOKEN_STRING:        {nil, nil, PREC_NONE},
		TOKEN_NUMBER:        {number, nil, PREC_NONE},
		TOKEN_AND:           {nil, nil, PREC_NONE},
		TOKEN_CLASS:         {nil, nil, PREC_NONE},
		TOKEN_ELSE:          {nil, nil, PREC_NONE},
		TOKEN_FALSE:         {literal, nil, PREC_NONE},
		TOKEN_FOR:           {nil, nil, PREC_NONE},
		TOKEN_FUN:           {nil, nil, PREC_NONE},
		TOKEN_IF:            {nil, nil, PREC_NONE},
		TOKEN_NIL:           {literal, nil, PREC_NONE},
		TOKEN_OR:            {nil, nil, PREC_NONE},
		TOKEN_PRINT:         {nil, nil, PREC_NONE},
		TOKEN_RETURN:        {nil, nil, PREC_NONE},
		TOKEN_SUPER:         {nil, nil, PREC_NONE},
		TOKEN_THIS:          {nil, nil, PREC_NONE},
		TOKEN_TRUE:          {literal, nil, PREC_NONE},
		TOKEN_VAR:           {nil, nil, PREC_NONE},
		TOKEN_WHILE:         {nil, nil, PREC_NONE},
		TOKEN_ERROR:         {nil, nil, PREC_NONE},
		TOKEN_EOF:           {nil, nil, PREC_NONE},
	}
	return rules[kind]
}

var (
	p              Parser       = Parser{}
	compilingChunk *chunk.Chunk = chunk.NewChunk()
)

func currentChunk() *chunk.Chunk {
	return compilingChunk
}

func advance() {
	p.previous = p.current

	for {
		p.current = p.scanner.NextToken()
		if p.current.GetTokenType() != TOKEN_ERROR {
			break
		}
		errorAtCurrent(p.current.String())
	}
}

func errorAtCurrent(message string) {
	errorAt(p.current, message)
}

func err(message string) {
	errorAt(p.previous, message)
}

func errorAt(token antlr.Token, message string) {
	if p.panicMode {
		return
	}
	p.panicMode = true
	fmt.Printf("[line %d] error", token.GetLine())

	if token.GetTokenType() == TOKEN_EOF {
		fmt.Printf("[line %d] error", token.GetLine())
	} else if token.GetTokenType() == TOKEN_ERROR {
		// nothing
	} else {
		fmt.Printf(" at '%s' %d", token.String(), token.GetStart())
	}

	fmt.Printf(": %s\n", message)
	p.hadError = true
}

func consume(kind int, message string) {
	if p.current.GetTokenType() == kind {
		advance()
		return
	}

	errorAtCurrent(message)
}

func emitByte(_byte byte) {
	currentChunk().WriteChunk(_byte, p.previous.GetLine())
}

func emitBytes(byte1, byte2 byte) {
	emitByte(byte1)
	emitByte(byte2)
}

func emitReturn() {
	emitByte(byte(chunk.OpReturn))
}

func emitConstant(value values.Value) {
	emitBytes(byte(chunk.OpConstant), makeConstant(value))
}

func endCompiler() {
	emitReturn()
	debug := true
	if debug {
		if !p.hadError {
			compilingChunk.DisassembleChunk("code")
		}
	}
}

func makeConstant(value values.Value) byte {
	constant := currentChunk().AddConstant(value)
	if constant > UINT8_MAX {
		err("too many constants in one chunk")
		return 0
	}
	return byte(constant)
}

func literal() {
	switch p.previous.GetTokenType() {
	case TOKEN_FALSE:
		emitByte(byte(chunk.OpFalse))
	case TOKEN_TRUE:
		emitByte(byte(chunk.OpTrue))
	case TOKEN_NIL:
		emitByte(byte(chunk.OpNil))
	}
}

func number() {
	value := float64(p.previous.GetStart())
	emitConstant(values.NewValue(values.VAL_NUMBER, value))
}

func grouping() {
	expression()
	consume(TOKEN_RIGHT_PAREN, "Expect ')' after expression")
}

func unary() {
	operatorType := p.previous.GetTokenType()

	parsePrecedence(PREC_UNARY)

	switch operatorType {
	case TOKEN_BANG:
		emitByte(byte(chunk.OpNot))
	case TOKEN_MINUS:
		emitByte(byte(chunk.OpNegate))
	default:
		return
	}
}

func binary() {
	operatorType := p.previous.GetTokenType()
	rule := getRule(operatorType)
	parsePrecedence(rule.precedence + 1)

	switch operatorType {
	case TOKEN_PLUS:
		emitByte(byte(chunk.OpAdd))
	case TOKEN_MINUS:
		emitByte(byte(chunk.OpSub))
	case TOKEN_STAR:
		emitByte(byte(chunk.OpMultiply))
	case TOKEN_SLASH:
		emitByte(byte(chunk.OpDivide))
	case TOKEN_BANG_EQUAL:
		emitBytes(byte(chunk.OpEqual), byte(chunk.OpNot))
	case TOKEN_EQUAL_EQUAL:
		emitByte(byte(chunk.OpEqual))
	case TOKEN_GREATER:
		emitByte(byte(chunk.OpGreater))
	case TOKEN_GREATER_EQUAL:
		emitBytes(byte(chunk.OpLess), byte(chunk.OpNot))
	case TOKEN_LESS:
		emitByte(byte(chunk.OpLess))
	case TOKEN_LESS_EQUAL:
		emitBytes(byte(chunk.OpGreater), byte(chunk.OpNot))
	default:
		return
	}
}

func parsePrecedence(precedence Precedence) {
	advance()
	prefixRule := getRule(p.previous.GetTokenType()).prefix
	if prefixRule == nil {
		err("expected expression")
		return
	}
	prefixRule()

	for precedence <= getRule(p.current.GetTokenType()).precedence {
		advance()
		infixRule := getRule(p.previous.GetTokenType()).infix
		infixRule()
	}
}

func expression() {
	parsePrecedence(PREC_ASSIGNMENT)
}

func Compile(source *antlr.InputStream, chunk *chunk.Chunk) bool {
	scanner := NewLexer(antlr.CharStream(source))
	p.scanner = scanner
	compilingChunk = chunk
	p.hadError = false
	p.panicMode = false
	advance()
	expression()

	consume(p.scanner.EmitEOF().GetTokenType(), "Expect end of expression.")
	endCompiler()
	return !p.hadError
}
