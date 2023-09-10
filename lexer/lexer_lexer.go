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

type LexerLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LexerLexerLexerStaticData struct {
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

func lexerlexerLexerInit() {
	staticData := &LexerLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'program'", "'class'", "'fun'", "'var'", "'for'", "'if'", "'else'",
		"'print'", "'return'", "'while'", "'true'", "'false'", "'nil'", "'this'",
		"'super'", "'or'", "'and'", "'!'", "'-'", "'+'", "'/'", "'*'", "'='",
		"'!='", "'=='", "'>'", "'>='", "'<'", "'<='", "'('", "')'", "'{'", "'}'",
		"','", "';'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "PROGRAM", "CLASS", "FUN", "VAR", "FOR", "IF", "ELSE", "PRINT",
		"RETURN", "WHILE", "TRUE", "FALSE", "NIL", "THIS", "SUPER", "OR", "AND",
		"NOT", "MINUS", "PLUS", "SLASH", "STAR", "EQUAL", "NOT_EQUAL", "EQUAL_EQUAL",
		"GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL", "LEFT_PAREN", "RIGHT_PAREN",
		"LEFT_BRACE", "RIGHT_BRACE", "COMMA", "SEMICOLON", "DOT", "NUMBER",
		"STRING", "IDENTIFIER", "WS",
	}
	staticData.RuleNames = []string{
		"PROGRAM", "CLASS", "FUN", "VAR", "FOR", "IF", "ELSE", "PRINT", "RETURN",
		"WHILE", "TRUE", "FALSE", "NIL", "THIS", "SUPER", "OR", "AND", "NOT",
		"MINUS", "PLUS", "SLASH", "STAR", "EQUAL", "NOT_EQUAL", "EQUAL_EQUAL",
		"GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL", "LEFT_PAREN", "RIGHT_PAREN",
		"LEFT_BRACE", "RIGHT_BRACE", "COMMA", "SEMICOLON", "DOT", "NUMBER",
		"STRING", "IDENTIFIER", "WS", "DIGIT", "ALPHA",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 40, 251, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 4, 36, 215, 8, 36,
		11, 36, 12, 36, 216, 1, 36, 1, 36, 4, 36, 221, 8, 36, 11, 36, 12, 36, 222,
		3, 36, 225, 8, 36, 1, 37, 1, 37, 5, 37, 229, 8, 37, 10, 37, 12, 37, 232,
		9, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 5, 38, 239, 8, 38, 10, 38, 12,
		38, 242, 9, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41,
		0, 0, 42, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 0, 83, 0, 1, 0, 4, 1, 0, 34, 34, 3, 0,
		9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 254,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0,
		0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1,
		0, 0, 0, 0, 79, 1, 0, 0, 0, 1, 85, 1, 0, 0, 0, 3, 93, 1, 0, 0, 0, 5, 99,
		1, 0, 0, 0, 7, 103, 1, 0, 0, 0, 9, 107, 1, 0, 0, 0, 11, 111, 1, 0, 0, 0,
		13, 114, 1, 0, 0, 0, 15, 119, 1, 0, 0, 0, 17, 125, 1, 0, 0, 0, 19, 132,
		1, 0, 0, 0, 21, 138, 1, 0, 0, 0, 23, 143, 1, 0, 0, 0, 25, 149, 1, 0, 0,
		0, 27, 153, 1, 0, 0, 0, 29, 158, 1, 0, 0, 0, 31, 164, 1, 0, 0, 0, 33, 167,
		1, 0, 0, 0, 35, 171, 1, 0, 0, 0, 37, 173, 1, 0, 0, 0, 39, 175, 1, 0, 0,
		0, 41, 177, 1, 0, 0, 0, 43, 179, 1, 0, 0, 0, 45, 181, 1, 0, 0, 0, 47, 183,
		1, 0, 0, 0, 49, 186, 1, 0, 0, 0, 51, 189, 1, 0, 0, 0, 53, 191, 1, 0, 0,
		0, 55, 194, 1, 0, 0, 0, 57, 196, 1, 0, 0, 0, 59, 199, 1, 0, 0, 0, 61, 201,
		1, 0, 0, 0, 63, 203, 1, 0, 0, 0, 65, 205, 1, 0, 0, 0, 67, 207, 1, 0, 0,
		0, 69, 209, 1, 0, 0, 0, 71, 211, 1, 0, 0, 0, 73, 214, 1, 0, 0, 0, 75, 226,
		1, 0, 0, 0, 77, 235, 1, 0, 0, 0, 79, 243, 1, 0, 0, 0, 81, 247, 1, 0, 0,
		0, 83, 249, 1, 0, 0, 0, 85, 86, 5, 112, 0, 0, 86, 87, 5, 114, 0, 0, 87,
		88, 5, 111, 0, 0, 88, 89, 5, 103, 0, 0, 89, 90, 5, 114, 0, 0, 90, 91, 5,
		97, 0, 0, 91, 92, 5, 109, 0, 0, 92, 2, 1, 0, 0, 0, 93, 94, 5, 99, 0, 0,
		94, 95, 5, 108, 0, 0, 95, 96, 5, 97, 0, 0, 96, 97, 5, 115, 0, 0, 97, 98,
		5, 115, 0, 0, 98, 4, 1, 0, 0, 0, 99, 100, 5, 102, 0, 0, 100, 101, 5, 117,
		0, 0, 101, 102, 5, 110, 0, 0, 102, 6, 1, 0, 0, 0, 103, 104, 5, 118, 0,
		0, 104, 105, 5, 97, 0, 0, 105, 106, 5, 114, 0, 0, 106, 8, 1, 0, 0, 0, 107,
		108, 5, 102, 0, 0, 108, 109, 5, 111, 0, 0, 109, 110, 5, 114, 0, 0, 110,
		10, 1, 0, 0, 0, 111, 112, 5, 105, 0, 0, 112, 113, 5, 102, 0, 0, 113, 12,
		1, 0, 0, 0, 114, 115, 5, 101, 0, 0, 115, 116, 5, 108, 0, 0, 116, 117, 5,
		115, 0, 0, 117, 118, 5, 101, 0, 0, 118, 14, 1, 0, 0, 0, 119, 120, 5, 112,
		0, 0, 120, 121, 5, 114, 0, 0, 121, 122, 5, 105, 0, 0, 122, 123, 5, 110,
		0, 0, 123, 124, 5, 116, 0, 0, 124, 16, 1, 0, 0, 0, 125, 126, 5, 114, 0,
		0, 126, 127, 5, 101, 0, 0, 127, 128, 5, 116, 0, 0, 128, 129, 5, 117, 0,
		0, 129, 130, 5, 114, 0, 0, 130, 131, 5, 110, 0, 0, 131, 18, 1, 0, 0, 0,
		132, 133, 5, 119, 0, 0, 133, 134, 5, 104, 0, 0, 134, 135, 5, 105, 0, 0,
		135, 136, 5, 108, 0, 0, 136, 137, 5, 101, 0, 0, 137, 20, 1, 0, 0, 0, 138,
		139, 5, 116, 0, 0, 139, 140, 5, 114, 0, 0, 140, 141, 5, 117, 0, 0, 141,
		142, 5, 101, 0, 0, 142, 22, 1, 0, 0, 0, 143, 144, 5, 102, 0, 0, 144, 145,
		5, 97, 0, 0, 145, 146, 5, 108, 0, 0, 146, 147, 5, 115, 0, 0, 147, 148,
		5, 101, 0, 0, 148, 24, 1, 0, 0, 0, 149, 150, 5, 110, 0, 0, 150, 151, 5,
		105, 0, 0, 151, 152, 5, 108, 0, 0, 152, 26, 1, 0, 0, 0, 153, 154, 5, 116,
		0, 0, 154, 155, 5, 104, 0, 0, 155, 156, 5, 105, 0, 0, 156, 157, 5, 115,
		0, 0, 157, 28, 1, 0, 0, 0, 158, 159, 5, 115, 0, 0, 159, 160, 5, 117, 0,
		0, 160, 161, 5, 112, 0, 0, 161, 162, 5, 101, 0, 0, 162, 163, 5, 114, 0,
		0, 163, 30, 1, 0, 0, 0, 164, 165, 5, 111, 0, 0, 165, 166, 5, 114, 0, 0,
		166, 32, 1, 0, 0, 0, 167, 168, 5, 97, 0, 0, 168, 169, 5, 110, 0, 0, 169,
		170, 5, 100, 0, 0, 170, 34, 1, 0, 0, 0, 171, 172, 5, 33, 0, 0, 172, 36,
		1, 0, 0, 0, 173, 174, 5, 45, 0, 0, 174, 38, 1, 0, 0, 0, 175, 176, 5, 43,
		0, 0, 176, 40, 1, 0, 0, 0, 177, 178, 5, 47, 0, 0, 178, 42, 1, 0, 0, 0,
		179, 180, 5, 42, 0, 0, 180, 44, 1, 0, 0, 0, 181, 182, 5, 61, 0, 0, 182,
		46, 1, 0, 0, 0, 183, 184, 5, 33, 0, 0, 184, 185, 5, 61, 0, 0, 185, 48,
		1, 0, 0, 0, 186, 187, 5, 61, 0, 0, 187, 188, 5, 61, 0, 0, 188, 50, 1, 0,
		0, 0, 189, 190, 5, 62, 0, 0, 190, 52, 1, 0, 0, 0, 191, 192, 5, 62, 0, 0,
		192, 193, 5, 61, 0, 0, 193, 54, 1, 0, 0, 0, 194, 195, 5, 60, 0, 0, 195,
		56, 1, 0, 0, 0, 196, 197, 5, 60, 0, 0, 197, 198, 5, 61, 0, 0, 198, 58,
		1, 0, 0, 0, 199, 200, 5, 40, 0, 0, 200, 60, 1, 0, 0, 0, 201, 202, 5, 41,
		0, 0, 202, 62, 1, 0, 0, 0, 203, 204, 5, 123, 0, 0, 204, 64, 1, 0, 0, 0,
		205, 206, 5, 125, 0, 0, 206, 66, 1, 0, 0, 0, 207, 208, 5, 44, 0, 0, 208,
		68, 1, 0, 0, 0, 209, 210, 5, 59, 0, 0, 210, 70, 1, 0, 0, 0, 211, 212, 5,
		46, 0, 0, 212, 72, 1, 0, 0, 0, 213, 215, 3, 81, 40, 0, 214, 213, 1, 0,
		0, 0, 215, 216, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0,
		217, 224, 1, 0, 0, 0, 218, 220, 5, 46, 0, 0, 219, 221, 3, 81, 40, 0, 220,
		219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 223,
		1, 0, 0, 0, 223, 225, 1, 0, 0, 0, 224, 218, 1, 0, 0, 0, 224, 225, 1, 0,
		0, 0, 225, 74, 1, 0, 0, 0, 226, 230, 5, 34, 0, 0, 227, 229, 8, 0, 0, 0,
		228, 227, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230,
		231, 1, 0, 0, 0, 231, 233, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233, 234,
		5, 34, 0, 0, 234, 76, 1, 0, 0, 0, 235, 240, 3, 83, 41, 0, 236, 239, 3,
		83, 41, 0, 237, 239, 3, 81, 40, 0, 238, 236, 1, 0, 0, 0, 238, 237, 1, 0,
		0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0,
		241, 78, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 243, 244, 7, 1, 0, 0, 244, 245,
		1, 0, 0, 0, 245, 246, 6, 39, 0, 0, 246, 80, 1, 0, 0, 0, 247, 248, 7, 2,
		0, 0, 248, 82, 1, 0, 0, 0, 249, 250, 7, 3, 0, 0, 250, 84, 1, 0, 0, 0, 7,
		0, 216, 222, 224, 230, 238, 240, 1, 6, 0, 0,
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

// LexerLexerInit initializes any static state used to implement LexerLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLexerLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LexerLexerInit() {
	staticData := &LexerLexerLexerStaticData
	staticData.once.Do(lexerlexerLexerInit)
}

// NewLexerLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLexerLexer(input antlr.CharStream) *LexerLexer {
	LexerLexerInit()
	l := new(LexerLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LexerLexerLexerStaticData
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

// LexerLexer tokens.
const (
	LexerLexerPROGRAM       = 1
	LexerLexerCLASS         = 2
	LexerLexerFUN           = 3
	LexerLexerVAR           = 4
	LexerLexerFOR           = 5
	LexerLexerIF            = 6
	LexerLexerELSE          = 7
	LexerLexerPRINT         = 8
	LexerLexerRETURN        = 9
	LexerLexerWHILE         = 10
	LexerLexerTRUE          = 11
	LexerLexerFALSE         = 12
	LexerLexerNIL           = 13
	LexerLexerTHIS          = 14
	LexerLexerSUPER         = 15
	LexerLexerOR            = 16
	LexerLexerAND           = 17
	LexerLexerNOT           = 18
	LexerLexerMINUS         = 19
	LexerLexerPLUS          = 20
	LexerLexerSLASH         = 21
	LexerLexerSTAR          = 22
	LexerLexerEQUAL         = 23
	LexerLexerNOT_EQUAL     = 24
	LexerLexerEQUAL_EQUAL   = 25
	LexerLexerGREATER       = 26
	LexerLexerGREATER_EQUAL = 27
	LexerLexerLESS          = 28
	LexerLexerLESS_EQUAL    = 29
	LexerLexerLEFT_PAREN    = 30
	LexerLexerRIGHT_PAREN   = 31
	LexerLexerLEFT_BRACE    = 32
	LexerLexerRIGHT_BRACE   = 33
	LexerLexerCOMMA         = 34
	LexerLexerSEMICOLON     = 35
	LexerLexerDOT           = 36
	LexerLexerNUMBER        = 37
	LexerLexerSTRING        = 38
	LexerLexerIDENTIFIER    = 39
	LexerLexerWS            = 40
)
