package compiler

import (
	"fmt"
	"math"
	"strconv"

	"github.com/VictorMilhomem/golox/cmd/chunk"
	"github.com/VictorMilhomem/golox/cmd/values"
)

const UINT8_MAX = math.MaxUint8

type Parser struct {
	current   Token
	previous  Token
	hadError  bool
	panicMode bool
	scanner   *Scanner
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

func getRule(kind TokenType) ParseRule {
	rules := map[TokenType]ParseRule{
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
		TOKEN_BANG:          {nil, nil, PREC_NONE},
		TOKEN_BANG_EQUAL:    {nil, nil, PREC_NONE},
		TOKEN_EQUAL:         {nil, nil, PREC_NONE},
		TOKEN_EQUAL_EQUAL:   {nil, nil, PREC_NONE},
		TOKEN_GREATER:       {nil, nil, PREC_NONE},
		TOKEN_GREATER_EQUAL: {nil, nil, PREC_NONE},
		TOKEN_LESS:          {nil, nil, PREC_NONE},
		TOKEN_LESS_EQUAL:    {nil, nil, PREC_NONE},
		TOKEN_IDENTIFIER:    {nil, nil, PREC_NONE},
		TOKEN_STRING:        {nil, nil, PREC_NONE},
		TOKEN_NUMBER:        {number, nil, PREC_NONE},
		TOKEN_AND:           {nil, nil, PREC_NONE},
		TOKEN_CLASS:         {nil, nil, PREC_NONE},
		TOKEN_ELSE:          {nil, nil, PREC_NONE},
		TOKEN_FALSE:         {nil, nil, PREC_NONE},
		TOKEN_FOR:           {nil, nil, PREC_NONE},
		TOKEN_FUN:           {nil, nil, PREC_NONE},
		TOKEN_IF:            {nil, nil, PREC_NONE},
		TOKEN_NIL:           {nil, nil, PREC_NONE},
		TOKEN_OR:            {nil, nil, PREC_NONE},
		TOKEN_PRINT:         {nil, nil, PREC_NONE},
		TOKEN_RETURN:        {nil, nil, PREC_NONE},
		TOKEN_SUPER:         {nil, nil, PREC_NONE},
		TOKEN_THIS:          {nil, nil, PREC_NONE},
		TOKEN_TRUE:          {nil, nil, PREC_NONE},
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
		p.current = p.scanner.scanToken()
		if p.current.kind != TOKEN_ERROR {
			break
		}
		errorAtCurrent(string(p.current.start))
	}
}

func errorAtCurrent(message string) {
	errorAt(p.current, message)
}

func err(message string) {
	errorAt(p.previous, message)
}

func errorAt(token Token, message string) {
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
		fmt.Printf(" at '%.*s'", token.length, token.start)
	}

	fmt.Printf(": %s\n", message)
	p.hadError = true
}

func consume(kind TokenType, message string) {
	if p.current.kind == kind {
		advance()
		return
	}

	errorAtCurrent(message)
}

func emitByte(_byte byte) {
	currentChunk().WriteChunk(_byte, p.previous.line)
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

func number() {
	value, err := strconv.ParseFloat(p.previous.start, 64)
	if err != nil {
		fmt.Errorf("error parsing float")
		return
	}
	emitConstant(values.Value(value))
}

func grouping() {
	expression()
	consume(TOKEN_RIGHT_PAREN, "Expect ')' after expression")
}

func unary() {
	operatorType := p.previous.kind

	parsePrecedence(PREC_UNARY)

	switch operatorType {
	case TOKEN_MINUS:
		emitByte(byte(chunk.OpNegate))
	default:
		return
	}
}

func binary() {
	operatorType := p.previous.kind
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
	default:
		return
	}
}

func parsePrecedence(precedence Precedence) {
	advance()
	prefixRule := getRule(p.previous.kind).prefix
	if prefixRule == nil {
		err("expected expression")
		return
	}
	prefixRule()

	for precedence <= getRule(p.current.kind).precedence {
		advance()
		infixRule := getRule(p.previous.kind).infix
		infixRule()
	}
}

func expression() {
	parsePrecedence(PREC_ASSIGNMENT)
}

func Compile(source string, chunk *chunk.Chunk) bool {
	scanner := NewScanner([]rune(source))
	p.scanner = scanner
	compilingChunk = chunk
	p.hadError = false
	p.panicMode = false
	advance()
	expression()
	consume(TOKEN_EOF, "Expect end of expression.")
	endCompiler()
	return !p.hadError
}
