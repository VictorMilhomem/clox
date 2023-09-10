// Code generated from ./cmd/compiler/Lexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lexerLexerInit() {
	staticData := &LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'{'", "'}'", "','", "'.'", "'-'", "'+'", "';'", "'/'",
		"'*'", "'!'", "'!='", "'='", "'=='", "'>'", "'>='", "'<'", "'<='", "",
		"", "", "'and'", "'class'", "'else'", "'false'", "'for'", "'fun'", "'if'",
		"'nil'", "'or'", "'print'", "'return'", "'super'", "'this'", "'true'",
		"'var'", "'while'",
	}
	staticData.SymbolicNames = []string{
		"", "TOKEN_LEFT_PAREN", "TOKEN_RIGHT_PAREN", "TOKEN_LEFT_BRACE", "TOKEN_RIGHT_BRACE",
		"TOKEN_COMMA", "TOKEN_DOT", "TOKEN_MINUS", "TOKEN_PLUS", "TOKEN_SEMICOLON",
		"TOKEN_SLASH", "TOKEN_STAR", "TOKEN_BANG", "TOKEN_BANG_EQUAL", "TOKEN_EQUAL",
		"TOKEN_EQUAL_EQUAL", "TOKEN_GREATER", "TOKEN_GREATER_EQUAL", "TOKEN_LESS",
		"TOKEN_LESS_EQUAL", "TOKEN_IDENTIFIER", "TOKEN_STRING", "TOKEN_NUMBER",
		"TOKEN_AND", "TOKEN_CLASS", "TOKEN_ELSE", "TOKEN_FALSE", "TOKEN_FOR",
		"TOKEN_FUN", "TOKEN_IF", "TOKEN_NIL", "TOKEN_OR", "TOKEN_PRINT", "TOKEN_RETURN",
		"TOKEN_SUPER", "TOKEN_THIS", "TOKEN_TRUE", "TOKEN_VAR", "TOKEN_WHILE",
		"TOKEN_ERROR", "TOKEN_EOF",
	}
	staticData.RuleNames = []string{
		"TOKEN_LEFT_PAREN", "TOKEN_RIGHT_PAREN", "TOKEN_LEFT_BRACE", "TOKEN_RIGHT_BRACE",
		"TOKEN_COMMA", "TOKEN_DOT", "TOKEN_MINUS", "TOKEN_PLUS", "TOKEN_SEMICOLON",
		"TOKEN_SLASH", "TOKEN_STAR", "TOKEN_BANG", "TOKEN_BANG_EQUAL", "TOKEN_EQUAL",
		"TOKEN_EQUAL_EQUAL", "TOKEN_GREATER", "TOKEN_GREATER_EQUAL", "TOKEN_LESS",
		"TOKEN_LESS_EQUAL", "TOKEN_IDENTIFIER", "TOKEN_STRING", "TOKEN_NUMBER",
		"TOKEN_AND", "TOKEN_CLASS", "TOKEN_ELSE", "TOKEN_FALSE", "TOKEN_FOR",
		"TOKEN_FUN", "TOKEN_IF", "TOKEN_NIL", "TOKEN_OR", "TOKEN_PRINT", "TOKEN_RETURN",
		"TOKEN_SUPER", "TOKEN_THIS", "TOKEN_TRUE", "TOKEN_VAR", "TOKEN_WHILE",
		"TOKEN_ERROR", "TOKEN_EOF", "DIGIT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 40, 238, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 128,
		8, 19, 10, 19, 12, 19, 131, 9, 19, 1, 20, 1, 20, 5, 20, 135, 8, 20, 10,
		20, 12, 20, 138, 9, 20, 1, 20, 1, 20, 1, 21, 4, 21, 143, 8, 21, 11, 21,
		12, 21, 144, 1, 21, 1, 21, 4, 21, 149, 8, 21, 11, 21, 12, 21, 150, 3, 21,
		153, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39,
		1, 40, 1, 40, 0, 0, 41, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 0, 1, 0, 4, 3, 0, 65, 90, 95,
		95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 34, 34, 1, 0,
		48, 57, 241, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0,
		61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0,
		0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0,
		0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 1, 83, 1, 0, 0, 0, 3, 85, 1, 0,
		0, 0, 5, 87, 1, 0, 0, 0, 7, 89, 1, 0, 0, 0, 9, 91, 1, 0, 0, 0, 11, 93,
		1, 0, 0, 0, 13, 95, 1, 0, 0, 0, 15, 97, 1, 0, 0, 0, 17, 99, 1, 0, 0, 0,
		19, 101, 1, 0, 0, 0, 21, 103, 1, 0, 0, 0, 23, 105, 1, 0, 0, 0, 25, 107,
		1, 0, 0, 0, 27, 110, 1, 0, 0, 0, 29, 112, 1, 0, 0, 0, 31, 115, 1, 0, 0,
		0, 33, 117, 1, 0, 0, 0, 35, 120, 1, 0, 0, 0, 37, 122, 1, 0, 0, 0, 39, 125,
		1, 0, 0, 0, 41, 132, 1, 0, 0, 0, 43, 142, 1, 0, 0, 0, 45, 154, 1, 0, 0,
		0, 47, 158, 1, 0, 0, 0, 49, 164, 1, 0, 0, 0, 51, 169, 1, 0, 0, 0, 53, 175,
		1, 0, 0, 0, 55, 179, 1, 0, 0, 0, 57, 183, 1, 0, 0, 0, 59, 186, 1, 0, 0,
		0, 61, 190, 1, 0, 0, 0, 63, 193, 1, 0, 0, 0, 65, 199, 1, 0, 0, 0, 67, 206,
		1, 0, 0, 0, 69, 212, 1, 0, 0, 0, 71, 217, 1, 0, 0, 0, 73, 222, 1, 0, 0,
		0, 75, 226, 1, 0, 0, 0, 77, 232, 1, 0, 0, 0, 79, 234, 1, 0, 0, 0, 81, 236,
		1, 0, 0, 0, 83, 84, 5, 40, 0, 0, 84, 2, 1, 0, 0, 0, 85, 86, 5, 41, 0, 0,
		86, 4, 1, 0, 0, 0, 87, 88, 5, 123, 0, 0, 88, 6, 1, 0, 0, 0, 89, 90, 5,
		125, 0, 0, 90, 8, 1, 0, 0, 0, 91, 92, 5, 44, 0, 0, 92, 10, 1, 0, 0, 0,
		93, 94, 5, 46, 0, 0, 94, 12, 1, 0, 0, 0, 95, 96, 5, 45, 0, 0, 96, 14, 1,
		0, 0, 0, 97, 98, 5, 43, 0, 0, 98, 16, 1, 0, 0, 0, 99, 100, 5, 59, 0, 0,
		100, 18, 1, 0, 0, 0, 101, 102, 5, 47, 0, 0, 102, 20, 1, 0, 0, 0, 103, 104,
		5, 42, 0, 0, 104, 22, 1, 0, 0, 0, 105, 106, 5, 33, 0, 0, 106, 24, 1, 0,
		0, 0, 107, 108, 5, 33, 0, 0, 108, 109, 5, 61, 0, 0, 109, 26, 1, 0, 0, 0,
		110, 111, 5, 61, 0, 0, 111, 28, 1, 0, 0, 0, 112, 113, 5, 61, 0, 0, 113,
		114, 5, 61, 0, 0, 114, 30, 1, 0, 0, 0, 115, 116, 5, 62, 0, 0, 116, 32,
		1, 0, 0, 0, 117, 118, 5, 62, 0, 0, 118, 119, 5, 61, 0, 0, 119, 34, 1, 0,
		0, 0, 120, 121, 5, 60, 0, 0, 121, 36, 1, 0, 0, 0, 122, 123, 5, 60, 0, 0,
		123, 124, 5, 61, 0, 0, 124, 38, 1, 0, 0, 0, 125, 129, 7, 0, 0, 0, 126,
		128, 7, 1, 0, 0, 127, 126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127,
		1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 40, 1, 0, 0, 0, 131, 129, 1, 0,
		0, 0, 132, 136, 5, 34, 0, 0, 133, 135, 8, 2, 0, 0, 134, 133, 1, 0, 0, 0,
		135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137,
		139, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 140, 5, 34, 0, 0, 140, 42,
		1, 0, 0, 0, 141, 143, 3, 81, 40, 0, 142, 141, 1, 0, 0, 0, 143, 144, 1,
		0, 0, 0, 144, 142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 152, 1, 0, 0,
		0, 146, 148, 5, 46, 0, 0, 147, 149, 3, 81, 40, 0, 148, 147, 1, 0, 0, 0,
		149, 150, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151,
		153, 1, 0, 0, 0, 152, 146, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 44, 1,
		0, 0, 0, 154, 155, 5, 97, 0, 0, 155, 156, 5, 110, 0, 0, 156, 157, 5, 100,
		0, 0, 157, 46, 1, 0, 0, 0, 158, 159, 5, 99, 0, 0, 159, 160, 5, 108, 0,
		0, 160, 161, 5, 97, 0, 0, 161, 162, 5, 115, 0, 0, 162, 163, 5, 115, 0,
		0, 163, 48, 1, 0, 0, 0, 164, 165, 5, 101, 0, 0, 165, 166, 5, 108, 0, 0,
		166, 167, 5, 115, 0, 0, 167, 168, 5, 101, 0, 0, 168, 50, 1, 0, 0, 0, 169,
		170, 5, 102, 0, 0, 170, 171, 5, 97, 0, 0, 171, 172, 5, 108, 0, 0, 172,
		173, 5, 115, 0, 0, 173, 174, 5, 101, 0, 0, 174, 52, 1, 0, 0, 0, 175, 176,
		5, 102, 0, 0, 176, 177, 5, 111, 0, 0, 177, 178, 5, 114, 0, 0, 178, 54,
		1, 0, 0, 0, 179, 180, 5, 102, 0, 0, 180, 181, 5, 117, 0, 0, 181, 182, 5,
		110, 0, 0, 182, 56, 1, 0, 0, 0, 183, 184, 5, 105, 0, 0, 184, 185, 5, 102,
		0, 0, 185, 58, 1, 0, 0, 0, 186, 187, 5, 110, 0, 0, 187, 188, 5, 105, 0,
		0, 188, 189, 5, 108, 0, 0, 189, 60, 1, 0, 0, 0, 190, 191, 5, 111, 0, 0,
		191, 192, 5, 114, 0, 0, 192, 62, 1, 0, 0, 0, 193, 194, 5, 112, 0, 0, 194,
		195, 5, 114, 0, 0, 195, 196, 5, 105, 0, 0, 196, 197, 5, 110, 0, 0, 197,
		198, 5, 116, 0, 0, 198, 64, 1, 0, 0, 0, 199, 200, 5, 114, 0, 0, 200, 201,
		5, 101, 0, 0, 201, 202, 5, 116, 0, 0, 202, 203, 5, 117, 0, 0, 203, 204,
		5, 114, 0, 0, 204, 205, 5, 110, 0, 0, 205, 66, 1, 0, 0, 0, 206, 207, 5,
		115, 0, 0, 207, 208, 5, 117, 0, 0, 208, 209, 5, 112, 0, 0, 209, 210, 5,
		101, 0, 0, 210, 211, 5, 114, 0, 0, 211, 68, 1, 0, 0, 0, 212, 213, 5, 116,
		0, 0, 213, 214, 5, 104, 0, 0, 214, 215, 5, 105, 0, 0, 215, 216, 5, 115,
		0, 0, 216, 70, 1, 0, 0, 0, 217, 218, 5, 116, 0, 0, 218, 219, 5, 114, 0,
		0, 219, 220, 5, 117, 0, 0, 220, 221, 5, 101, 0, 0, 221, 72, 1, 0, 0, 0,
		222, 223, 5, 118, 0, 0, 223, 224, 5, 97, 0, 0, 224, 225, 5, 114, 0, 0,
		225, 74, 1, 0, 0, 0, 226, 227, 5, 119, 0, 0, 227, 228, 5, 104, 0, 0, 228,
		229, 5, 105, 0, 0, 229, 230, 5, 108, 0, 0, 230, 231, 5, 101, 0, 0, 231,
		76, 1, 0, 0, 0, 232, 233, 9, 0, 0, 0, 233, 78, 1, 0, 0, 0, 234, 235, 5,
		0, 0, 1, 235, 80, 1, 0, 0, 0, 236, 237, 7, 3, 0, 0, 237, 82, 1, 0, 0, 0,
		6, 0, 129, 136, 144, 150, 152, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LexerInit initializes any static state used to implement Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LexerInit() {
	staticData := &LexerLexerStaticData
	staticData.once.Do(lexerLexerInit)
}

// NewLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLexer(input antlr.CharStream) *Lexer {
	LexerInit()
	l := new(Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Lexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Lexer tokens.
const (
	LexerTOKEN_LEFT_PAREN    = 1
	LexerTOKEN_RIGHT_PAREN   = 2
	LexerTOKEN_LEFT_BRACE    = 3
	LexerTOKEN_RIGHT_BRACE   = 4
	LexerTOKEN_COMMA         = 5
	LexerTOKEN_DOT           = 6
	LexerTOKEN_MINUS         = 7
	LexerTOKEN_PLUS          = 8
	LexerTOKEN_SEMICOLON     = 9
	LexerTOKEN_SLASH         = 10
	LexerTOKEN_STAR          = 11
	LexerTOKEN_BANG          = 12
	LexerTOKEN_BANG_EQUAL    = 13
	LexerTOKEN_EQUAL         = 14
	LexerTOKEN_EQUAL_EQUAL   = 15
	LexerTOKEN_GREATER       = 16
	LexerTOKEN_GREATER_EQUAL = 17
	LexerTOKEN_LESS          = 18
	LexerTOKEN_LESS_EQUAL    = 19
	LexerTOKEN_IDENTIFIER    = 20
	LexerTOKEN_STRING        = 21
	LexerTOKEN_NUMBER        = 22
	LexerTOKEN_AND           = 23
	LexerTOKEN_CLASS         = 24
	LexerTOKEN_ELSE          = 25
	LexerTOKEN_FALSE         = 26
	LexerTOKEN_FOR           = 27
	LexerTOKEN_FUN           = 28
	LexerTOKEN_IF            = 29
	LexerTOKEN_NIL           = 30
	LexerTOKEN_OR            = 31
	LexerTOKEN_PRINT         = 32
	LexerTOKEN_RETURN        = 33
	LexerTOKEN_SUPER         = 34
	LexerTOKEN_THIS          = 35
	LexerTOKEN_TRUE          = 36
	LexerTOKEN_VAR           = 37
	LexerTOKEN_WHILE         = 38
	LexerTOKEN_ERROR         = 39
	LexerTOKEN_EOF           = 40
)
