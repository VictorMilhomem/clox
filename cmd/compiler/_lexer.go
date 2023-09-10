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
		"", "", "'('", "')'", "'{'", "'}'", "','", "'.'", "'-'", "'+'", "';'",
		"'/'", "'*'", "'!'", "'!='", "'='", "'=='", "'>'", "'>='", "'<'", "'<='",
		"'and'", "'class'", "'else'", "'false'", "'for'", "'fun'", "'if'", "'nil'",
		"'or'", "'print'", "'return'", "'super'", "'this'", "'true'", "'var'",
		"'while'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "TOKEN_LEFT_PAREN", "TOKEN_RIGHT_PAREN", "TOKEN_LEFT_BRACE",
		"TOKEN_RIGHT_BRACE", "TOKEN_COMMA", "TOKEN_DOT", "TOKEN_MINUS", "TOKEN_PLUS",
		"TOKEN_SEMICOLON", "TOKEN_SLASH", "TOKEN_STAR", "TOKEN_BANG", "TOKEN_BANG_EQUAL",
		"TOKEN_EQUAL", "TOKEN_EQUAL_EQUAL", "TOKEN_GREATER", "TOKEN_GREATER_EQUAL",
		"TOKEN_LESS", "TOKEN_LESS_EQUAL", "TOKEN_AND", "TOKEN_CLASS", "TOKEN_ELSE",
		"TOKEN_FALSE", "TOKEN_FOR", "TOKEN_FUN", "TOKEN_IF", "TOKEN_NIL", "TOKEN_OR",
		"TOKEN_PRINT", "TOKEN_RETURN", "TOKEN_SUPER", "TOKEN_THIS", "TOKEN_TRUE",
		"TOKEN_VAR", "TOKEN_WHILE", "TOKEN_STRING", "TOKEN_NUMBER", "TOKEN_IDENTIFIER",
		"TOKEN_EOF", "TOKEN_ERROR",
	}
	staticData.RuleNames = []string{
		"WS", "TOKEN_LEFT_PAREN", "TOKEN_RIGHT_PAREN", "TOKEN_LEFT_BRACE", "TOKEN_RIGHT_BRACE",
		"TOKEN_COMMA", "TOKEN_DOT", "TOKEN_MINUS", "TOKEN_PLUS", "TOKEN_SEMICOLON",
		"TOKEN_SLASH", "TOKEN_STAR", "TOKEN_BANG", "TOKEN_BANG_EQUAL", "TOKEN_EQUAL",
		"TOKEN_EQUAL_EQUAL", "TOKEN_GREATER", "TOKEN_GREATER_EQUAL", "TOKEN_LESS",
		"TOKEN_LESS_EQUAL", "TOKEN_AND", "TOKEN_CLASS", "TOKEN_ELSE", "TOKEN_FALSE",
		"TOKEN_FOR", "TOKEN_FUN", "TOKEN_IF", "TOKEN_NIL", "TOKEN_OR", "TOKEN_PRINT",
		"TOKEN_RETURN", "TOKEN_SUPER", "TOKEN_THIS", "TOKEN_TRUE", "TOKEN_VAR",
		"TOKEN_WHILE", "TOKEN_STRING", "TOKEN_NUMBER", "TOKEN_IDENTIFIER", "TOKEN_EOF",
		"TOKEN_ERROR", "DIGIT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 41, 249, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 1, 0, 4, 0, 87, 8, 0, 11, 0, 12, 0, 88, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 36, 1, 36, 5, 36, 215, 8, 36, 10, 36, 12, 36, 218, 9, 36, 1, 36, 1,
		36, 1, 37, 4, 37, 223, 8, 37, 11, 37, 12, 37, 224, 1, 37, 1, 37, 4, 37,
		229, 8, 37, 11, 37, 12, 37, 230, 3, 37, 233, 8, 37, 1, 38, 1, 38, 5, 38,
		237, 8, 38, 10, 38, 12, 38, 240, 9, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 41, 1, 41, 0, 0, 42, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49,
		25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67,
		34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 0, 1, 0,
		5, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 34, 34, 3, 0, 65, 90, 95, 95, 97,
		122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 253, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63,
		1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0,
		71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0,
		0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 1, 86, 1, 0, 0, 0, 3, 92, 1, 0, 0,
		0, 5, 94, 1, 0, 0, 0, 7, 96, 1, 0, 0, 0, 9, 98, 1, 0, 0, 0, 11, 100, 1,
		0, 0, 0, 13, 102, 1, 0, 0, 0, 15, 104, 1, 0, 0, 0, 17, 106, 1, 0, 0, 0,
		19, 108, 1, 0, 0, 0, 21, 110, 1, 0, 0, 0, 23, 112, 1, 0, 0, 0, 25, 114,
		1, 0, 0, 0, 27, 116, 1, 0, 0, 0, 29, 119, 1, 0, 0, 0, 31, 121, 1, 0, 0,
		0, 33, 124, 1, 0, 0, 0, 35, 126, 1, 0, 0, 0, 37, 129, 1, 0, 0, 0, 39, 131,
		1, 0, 0, 0, 41, 134, 1, 0, 0, 0, 43, 138, 1, 0, 0, 0, 45, 144, 1, 0, 0,
		0, 47, 149, 1, 0, 0, 0, 49, 155, 1, 0, 0, 0, 51, 159, 1, 0, 0, 0, 53, 163,
		1, 0, 0, 0, 55, 166, 1, 0, 0, 0, 57, 170, 1, 0, 0, 0, 59, 173, 1, 0, 0,
		0, 61, 179, 1, 0, 0, 0, 63, 186, 1, 0, 0, 0, 65, 192, 1, 0, 0, 0, 67, 197,
		1, 0, 0, 0, 69, 202, 1, 0, 0, 0, 71, 206, 1, 0, 0, 0, 73, 212, 1, 0, 0,
		0, 75, 222, 1, 0, 0, 0, 77, 234, 1, 0, 0, 0, 79, 241, 1, 0, 0, 0, 81, 243,
		1, 0, 0, 0, 83, 247, 1, 0, 0, 0, 85, 87, 7, 0, 0, 0, 86, 85, 1, 0, 0, 0,
		87, 88, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 1,
		0, 0, 0, 90, 91, 6, 0, 0, 0, 91, 2, 1, 0, 0, 0, 92, 93, 5, 40, 0, 0, 93,
		4, 1, 0, 0, 0, 94, 95, 5, 41, 0, 0, 95, 6, 1, 0, 0, 0, 96, 97, 5, 123,
		0, 0, 97, 8, 1, 0, 0, 0, 98, 99, 5, 125, 0, 0, 99, 10, 1, 0, 0, 0, 100,
		101, 5, 44, 0, 0, 101, 12, 1, 0, 0, 0, 102, 103, 5, 46, 0, 0, 103, 14,
		1, 0, 0, 0, 104, 105, 5, 45, 0, 0, 105, 16, 1, 0, 0, 0, 106, 107, 5, 43,
		0, 0, 107, 18, 1, 0, 0, 0, 108, 109, 5, 59, 0, 0, 109, 20, 1, 0, 0, 0,
		110, 111, 5, 47, 0, 0, 111, 22, 1, 0, 0, 0, 112, 113, 5, 42, 0, 0, 113,
		24, 1, 0, 0, 0, 114, 115, 5, 33, 0, 0, 115, 26, 1, 0, 0, 0, 116, 117, 5,
		33, 0, 0, 117, 118, 5, 61, 0, 0, 118, 28, 1, 0, 0, 0, 119, 120, 5, 61,
		0, 0, 120, 30, 1, 0, 0, 0, 121, 122, 5, 61, 0, 0, 122, 123, 5, 61, 0, 0,
		123, 32, 1, 0, 0, 0, 124, 125, 5, 62, 0, 0, 125, 34, 1, 0, 0, 0, 126, 127,
		5, 62, 0, 0, 127, 128, 5, 61, 0, 0, 128, 36, 1, 0, 0, 0, 129, 130, 5, 60,
		0, 0, 130, 38, 1, 0, 0, 0, 131, 132, 5, 60, 0, 0, 132, 133, 5, 61, 0, 0,
		133, 40, 1, 0, 0, 0, 134, 135, 5, 97, 0, 0, 135, 136, 5, 110, 0, 0, 136,
		137, 5, 100, 0, 0, 137, 42, 1, 0, 0, 0, 138, 139, 5, 99, 0, 0, 139, 140,
		5, 108, 0, 0, 140, 141, 5, 97, 0, 0, 141, 142, 5, 115, 0, 0, 142, 143,
		5, 115, 0, 0, 143, 44, 1, 0, 0, 0, 144, 145, 5, 101, 0, 0, 145, 146, 5,
		108, 0, 0, 146, 147, 5, 115, 0, 0, 147, 148, 5, 101, 0, 0, 148, 46, 1,
		0, 0, 0, 149, 150, 5, 102, 0, 0, 150, 151, 5, 97, 0, 0, 151, 152, 5, 108,
		0, 0, 152, 153, 5, 115, 0, 0, 153, 154, 5, 101, 0, 0, 154, 48, 1, 0, 0,
		0, 155, 156, 5, 102, 0, 0, 156, 157, 5, 111, 0, 0, 157, 158, 5, 114, 0,
		0, 158, 50, 1, 0, 0, 0, 159, 160, 5, 102, 0, 0, 160, 161, 5, 117, 0, 0,
		161, 162, 5, 110, 0, 0, 162, 52, 1, 0, 0, 0, 163, 164, 5, 105, 0, 0, 164,
		165, 5, 102, 0, 0, 165, 54, 1, 0, 0, 0, 166, 167, 5, 110, 0, 0, 167, 168,
		5, 105, 0, 0, 168, 169, 5, 108, 0, 0, 169, 56, 1, 0, 0, 0, 170, 171, 5,
		111, 0, 0, 171, 172, 5, 114, 0, 0, 172, 58, 1, 0, 0, 0, 173, 174, 5, 112,
		0, 0, 174, 175, 5, 114, 0, 0, 175, 176, 5, 105, 0, 0, 176, 177, 5, 110,
		0, 0, 177, 178, 5, 116, 0, 0, 178, 60, 1, 0, 0, 0, 179, 180, 5, 114, 0,
		0, 180, 181, 5, 101, 0, 0, 181, 182, 5, 116, 0, 0, 182, 183, 5, 117, 0,
		0, 183, 184, 5, 114, 0, 0, 184, 185, 5, 110, 0, 0, 185, 62, 1, 0, 0, 0,
		186, 187, 5, 115, 0, 0, 187, 188, 5, 117, 0, 0, 188, 189, 5, 112, 0, 0,
		189, 190, 5, 101, 0, 0, 190, 191, 5, 114, 0, 0, 191, 64, 1, 0, 0, 0, 192,
		193, 5, 116, 0, 0, 193, 194, 5, 104, 0, 0, 194, 195, 5, 105, 0, 0, 195,
		196, 5, 115, 0, 0, 196, 66, 1, 0, 0, 0, 197, 198, 5, 116, 0, 0, 198, 199,
		5, 114, 0, 0, 199, 200, 5, 117, 0, 0, 200, 201, 5, 101, 0, 0, 201, 68,
		1, 0, 0, 0, 202, 203, 5, 118, 0, 0, 203, 204, 5, 97, 0, 0, 204, 205, 5,
		114, 0, 0, 205, 70, 1, 0, 0, 0, 206, 207, 5, 119, 0, 0, 207, 208, 5, 104,
		0, 0, 208, 209, 5, 105, 0, 0, 209, 210, 5, 108, 0, 0, 210, 211, 5, 101,
		0, 0, 211, 72, 1, 0, 0, 0, 212, 216, 5, 34, 0, 0, 213, 215, 8, 1, 0, 0,
		214, 213, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216,
		217, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 220,
		5, 34, 0, 0, 220, 74, 1, 0, 0, 0, 221, 223, 3, 83, 41, 0, 222, 221, 1,
		0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0,
		0, 225, 232, 1, 0, 0, 0, 226, 228, 5, 46, 0, 0, 227, 229, 3, 83, 41, 0,
		228, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230,
		231, 1, 0, 0, 0, 231, 233, 1, 0, 0, 0, 232, 226, 1, 0, 0, 0, 232, 233,
		1, 0, 0, 0, 233, 76, 1, 0, 0, 0, 234, 238, 7, 2, 0, 0, 235, 237, 7, 3,
		0, 0, 236, 235, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0,
		238, 239, 1, 0, 0, 0, 239, 78, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 241, 242,
		5, 0, 0, 1, 242, 80, 1, 0, 0, 0, 243, 244, 9, 0, 0, 0, 244, 245, 1, 0,
		0, 0, 245, 246, 6, 40, 0, 0, 246, 82, 1, 0, 0, 0, 247, 248, 7, 4, 0, 0,
		248, 84, 1, 0, 0, 0, 7, 0, 88, 216, 224, 230, 232, 238, 1, 6, 0, 0,
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
	LexerWS                  = 1
	LexerTOKEN_LEFT_PAREN    = 2
	LexerTOKEN_RIGHT_PAREN   = 3
	LexerTOKEN_LEFT_BRACE    = 4
	LexerTOKEN_RIGHT_BRACE   = 5
	LexerTOKEN_COMMA         = 6
	LexerTOKEN_DOT           = 7
	LexerTOKEN_MINUS         = 8
	LexerTOKEN_PLUS          = 9
	LexerTOKEN_SEMICOLON     = 10
	LexerTOKEN_SLASH         = 11
	LexerTOKEN_STAR          = 12
	LexerTOKEN_BANG          = 13
	LexerTOKEN_BANG_EQUAL    = 14
	LexerTOKEN_EQUAL         = 15
	LexerTOKEN_EQUAL_EQUAL   = 16
	LexerTOKEN_GREATER       = 17
	LexerTOKEN_GREATER_EQUAL = 18
	LexerTOKEN_LESS          = 19
	LexerTOKEN_LESS_EQUAL    = 20
	LexerTOKEN_AND           = 21
	LexerTOKEN_CLASS         = 22
	LexerTOKEN_ELSE          = 23
	LexerTOKEN_FALSE         = 24
	LexerTOKEN_FOR           = 25
	LexerTOKEN_FUN           = 26
	LexerTOKEN_IF            = 27
	LexerTOKEN_NIL           = 28
	LexerTOKEN_OR            = 29
	LexerTOKEN_PRINT         = 30
	LexerTOKEN_RETURN        = 31
	LexerTOKEN_SUPER         = 32
	LexerTOKEN_THIS          = 33
	LexerTOKEN_TRUE          = 34
	LexerTOKEN_VAR           = 35
	LexerTOKEN_WHILE         = 36
	LexerTOKEN_STRING        = 37
	LexerTOKEN_NUMBER        = 38
	LexerTOKEN_IDENTIFIER    = 39
	LexerTOKEN_EOF           = 40
	LexerTOKEN_ERROR         = 41
)
