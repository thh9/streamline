// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type RiLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var rilexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rilexerLexerInit() {
	staticData := &rilexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'('", "')'", "'{'", "'}'", "'='", "';'", "'echo'", "'if'", "'for'",
		"'bool'", "'string'", "'float'", "'any'", "'func'", "'true'", "'false'",
		"", "", "", "", "", "'*'", "'/'", "'+'", "'-'", "'%'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "ECHO", "IF", "FOR", "BOOL", "STRING", "FLOAT",
		"ANY", "FUNC", "TRUE", "FALSE", "ID", "INT", "NEWLINE", "WS", "DECIMAL_FLOAT_LIT",
		"MUL", "DIV", "ADD", "SUB", "MOD",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "ECHO", "IF", "FOR",
		"BOOL", "STRING", "FLOAT", "ANY", "FUNC", "TRUE", "FALSE", "ID", "INT",
		"NEWLINE", "WS", "DECIMALS", "DECIMAL_FLOAT_LIT", "MUL", "DIV", "ADD",
		"SUB", "MOD",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 164, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 4, 16, 119,
		8, 16, 11, 16, 12, 16, 120, 1, 17, 4, 17, 124, 8, 17, 11, 17, 12, 17, 125,
		1, 18, 3, 18, 129, 8, 18, 1, 18, 1, 18, 1, 19, 4, 19, 134, 8, 19, 11, 19,
		12, 19, 135, 1, 19, 1, 19, 1, 20, 1, 20, 3, 20, 142, 8, 20, 1, 20, 5, 20,
		145, 8, 20, 10, 20, 12, 20, 148, 9, 20, 1, 21, 1, 21, 1, 21, 3, 21, 153,
		8, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 0, 0, 27, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 39, 20, 41, 0, 43, 21, 45, 22, 47, 23, 49, 24, 51, 25, 53, 26,
		1, 0, 3, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 2, 0, 9, 9, 32, 32, 169,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 1, 55, 1, 0, 0, 0,
		3, 57, 1, 0, 0, 0, 5, 59, 1, 0, 0, 0, 7, 61, 1, 0, 0, 0, 9, 63, 1, 0, 0,
		0, 11, 65, 1, 0, 0, 0, 13, 67, 1, 0, 0, 0, 15, 72, 1, 0, 0, 0, 17, 75,
		1, 0, 0, 0, 19, 79, 1, 0, 0, 0, 21, 84, 1, 0, 0, 0, 23, 91, 1, 0, 0, 0,
		25, 97, 1, 0, 0, 0, 27, 101, 1, 0, 0, 0, 29, 106, 1, 0, 0, 0, 31, 111,
		1, 0, 0, 0, 33, 118, 1, 0, 0, 0, 35, 123, 1, 0, 0, 0, 37, 128, 1, 0, 0,
		0, 39, 133, 1, 0, 0, 0, 41, 139, 1, 0, 0, 0, 43, 149, 1, 0, 0, 0, 45, 154,
		1, 0, 0, 0, 47, 156, 1, 0, 0, 0, 49, 158, 1, 0, 0, 0, 51, 160, 1, 0, 0,
		0, 53, 162, 1, 0, 0, 0, 55, 56, 5, 40, 0, 0, 56, 2, 1, 0, 0, 0, 57, 58,
		5, 41, 0, 0, 58, 4, 1, 0, 0, 0, 59, 60, 5, 123, 0, 0, 60, 6, 1, 0, 0, 0,
		61, 62, 5, 125, 0, 0, 62, 8, 1, 0, 0, 0, 63, 64, 5, 61, 0, 0, 64, 10, 1,
		0, 0, 0, 65, 66, 5, 59, 0, 0, 66, 12, 1, 0, 0, 0, 67, 68, 5, 101, 0, 0,
		68, 69, 5, 99, 0, 0, 69, 70, 5, 104, 0, 0, 70, 71, 5, 111, 0, 0, 71, 14,
		1, 0, 0, 0, 72, 73, 5, 105, 0, 0, 73, 74, 5, 102, 0, 0, 74, 16, 1, 0, 0,
		0, 75, 76, 5, 102, 0, 0, 76, 77, 5, 111, 0, 0, 77, 78, 5, 114, 0, 0, 78,
		18, 1, 0, 0, 0, 79, 80, 5, 98, 0, 0, 80, 81, 5, 111, 0, 0, 81, 82, 5, 111,
		0, 0, 82, 83, 5, 108, 0, 0, 83, 20, 1, 0, 0, 0, 84, 85, 5, 115, 0, 0, 85,
		86, 5, 116, 0, 0, 86, 87, 5, 114, 0, 0, 87, 88, 5, 105, 0, 0, 88, 89, 5,
		110, 0, 0, 89, 90, 5, 103, 0, 0, 90, 22, 1, 0, 0, 0, 91, 92, 5, 102, 0,
		0, 92, 93, 5, 108, 0, 0, 93, 94, 5, 111, 0, 0, 94, 95, 5, 97, 0, 0, 95,
		96, 5, 116, 0, 0, 96, 24, 1, 0, 0, 0, 97, 98, 5, 97, 0, 0, 98, 99, 5, 110,
		0, 0, 99, 100, 5, 121, 0, 0, 100, 26, 1, 0, 0, 0, 101, 102, 5, 102, 0,
		0, 102, 103, 5, 117, 0, 0, 103, 104, 5, 110, 0, 0, 104, 105, 5, 99, 0,
		0, 105, 28, 1, 0, 0, 0, 106, 107, 5, 116, 0, 0, 107, 108, 5, 114, 0, 0,
		108, 109, 5, 117, 0, 0, 109, 110, 5, 101, 0, 0, 110, 30, 1, 0, 0, 0, 111,
		112, 5, 102, 0, 0, 112, 113, 5, 97, 0, 0, 113, 114, 5, 108, 0, 0, 114,
		115, 5, 115, 0, 0, 115, 116, 5, 101, 0, 0, 116, 32, 1, 0, 0, 0, 117, 119,
		7, 0, 0, 0, 118, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 118, 1, 0,
		0, 0, 120, 121, 1, 0, 0, 0, 121, 34, 1, 0, 0, 0, 122, 124, 7, 1, 0, 0,
		123, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125,
		126, 1, 0, 0, 0, 126, 36, 1, 0, 0, 0, 127, 129, 5, 13, 0, 0, 128, 127,
		1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 5, 10,
		0, 0, 131, 38, 1, 0, 0, 0, 132, 134, 7, 2, 0, 0, 133, 132, 1, 0, 0, 0,
		134, 135, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136,
		137, 1, 0, 0, 0, 137, 138, 6, 19, 0, 0, 138, 40, 1, 0, 0, 0, 139, 146,
		7, 1, 0, 0, 140, 142, 5, 95, 0, 0, 141, 140, 1, 0, 0, 0, 141, 142, 1, 0,
		0, 0, 142, 143, 1, 0, 0, 0, 143, 145, 7, 1, 0, 0, 144, 141, 1, 0, 0, 0,
		145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147,
		42, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 150, 3, 41, 20, 0, 150, 152,
		5, 46, 0, 0, 151, 153, 3, 41, 20, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1,
		0, 0, 0, 153, 44, 1, 0, 0, 0, 154, 155, 5, 42, 0, 0, 155, 46, 1, 0, 0,
		0, 156, 157, 5, 47, 0, 0, 157, 48, 1, 0, 0, 0, 158, 159, 5, 43, 0, 0, 159,
		50, 1, 0, 0, 0, 160, 161, 5, 45, 0, 0, 161, 52, 1, 0, 0, 0, 162, 163, 5,
		37, 0, 0, 163, 54, 1, 0, 0, 0, 8, 0, 120, 125, 128, 135, 141, 146, 152,
		1, 6, 0, 0,
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

// RiLexerInit initializes any static state used to implement RiLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRiLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RiLexerInit() {
	staticData := &rilexerLexerStaticData
	staticData.once.Do(rilexerLexerInit)
}

// NewRiLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRiLexer(input antlr.CharStream) *RiLexer {
	RiLexerInit()
	l := new(RiLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &rilexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Ri.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RiLexer tokens.
const (
	RiLexerT__0              = 1
	RiLexerT__1              = 2
	RiLexerT__2              = 3
	RiLexerT__3              = 4
	RiLexerT__4              = 5
	RiLexerT__5              = 6
	RiLexerECHO              = 7
	RiLexerIF                = 8
	RiLexerFOR               = 9
	RiLexerBOOL              = 10
	RiLexerSTRING            = 11
	RiLexerFLOAT             = 12
	RiLexerANY               = 13
	RiLexerFUNC              = 14
	RiLexerTRUE              = 15
	RiLexerFALSE             = 16
	RiLexerID                = 17
	RiLexerINT               = 18
	RiLexerNEWLINE           = 19
	RiLexerWS                = 20
	RiLexerDECIMAL_FLOAT_LIT = 21
	RiLexerMUL               = 22
	RiLexerDIV               = 23
	RiLexerADD               = 24
	RiLexerSUB               = 25
	RiLexerMOD               = 26
)
