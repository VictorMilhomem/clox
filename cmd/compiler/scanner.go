package compiler

import (
	"regexp"
	"unicode"
)

type TokenType int

const (
	// Single-character tokens.
	TOKEN_LEFT_PAREN TokenType = iota
	TOKEN_RIGHT_PAREN
	TOKEN_LEFT_BRACE
	TOKEN_RIGHT_BRACE
	TOKEN_COMMA
	TOKEN_DOT
	TOKEN_MINUS
	TOKEN_PLUS
	TOKEN_SEMICOLON
	TOKEN_SLASH
	TOKEN_STAR
	//	One or two character tokens.

	TOKEN_BANG
	TOKEN_BANG_EQUAL
	TOKEN_EQUAL
	TOKEN_EQUAL_EQUAL
	TOKEN_GREATER
	TOKEN_GREATER_EQUAL
	TOKEN_LESS
	TOKEN_LESS_EQUAL
	// Literals
	TOKEN_IDENTIFIER
	TOKEN_STRING
	TOKEN_NUMBER
	// Keywords
	TOKEN_AND
	TOKEN_CLASS
	TOKEN_ELSE
	TOKEN_FALSE
	TOKEN_FOR
	TOKEN_FUN
	TOKEN_IF
	TOKEN_NIL
	TOKEN_OR
	TOKEN_PRINT
	TOKEN_RETURN
	TOKEN_SUPER
	TOKEN_THIS
	TOKEN_TRUE
	TOKEN_VAR
	TOKEN_WHILE

	TOKEN_ERROR
	TOKEN_EOF
)

type Token struct {
	kind   TokenType
	start  string
	length int
	line   int
}

type Scanner struct {
	source  []rune
	start   rune
	current rune
	length  int
	at      int
	line    int
}

func NewToken(kind TokenType, start string, length int, line int) Token {
	return Token{
		kind:   kind,
		start:  start,
		length: length,
		line:   line,
	}
}

func NewScanner(source []rune) *Scanner {
	return &Scanner{
		source:  source,
		start:   source[0],
		current: source[0],
		length:  len(source),
		at:      0,
		line:    1,
	}
}

func (s *Scanner) scanToken() Token {
	s.skipWhiteSpace()
	s.start = s.current

	if s.isAtEnd() {
		return s.makeToken(TOKEN_EOF)
	}

	char := s.advance()
	if s.isAlpha(char) {
		return s.identifier()
	}
	if unicode.IsDigit(char) {
		return s.number()
	}

	switch char {
	case '(':
		return s.makeToken(TOKEN_LEFT_PAREN)
	case ')':
		return s.makeToken(TOKEN_RIGHT_PAREN)
	case '{':
		return s.makeToken(TOKEN_LEFT_BRACE)
	case '}':
		return s.makeToken(TOKEN_RIGHT_BRACE)
	case ';':
		return s.makeToken(TOKEN_SEMICOLON)
	case ',':
		return s.makeToken(TOKEN_COMMA)
	case '.':
		return s.makeToken(TOKEN_DOT)
	case '-':
		return s.makeToken(TOKEN_MINUS)
	case '+':
		return s.makeToken(TOKEN_PLUS)
	case '/':
		return s.makeToken(TOKEN_SLASH)
	case '*':
		return s.makeToken(TOKEN_STAR)
	case '!':
		if s.match('=') {
			return s.makeToken(TOKEN_BANG_EQUAL)
		}
		return s.makeToken(TOKEN_BANG)
	case '=':
		if s.match('=') {
			return s.makeToken(TOKEN_EQUAL_EQUAL)
		}
		return s.makeToken(TOKEN_EQUAL)
	case '<':
		if s.match('=') {
			return s.makeToken(TOKEN_LESS_EQUAL)
		}
		return s.makeToken(TOKEN_LESS)
	case '>':
		if s.match('=') {
			return s.makeToken(TOKEN_GREATER_EQUAL)
		}
		return s.makeToken(TOKEN_GREATER)
	case '"':
		return s.str()
	}

	return s.errorToken("unexpected character")
}

func (s *Scanner) isAtEnd() bool {
	return s.current == '\x00'
}

func (s *Scanner) makeToken(token TokenType) Token {
	return NewToken(token, string(s.start), int(s.current-s.start), s.line)
}

func (s *Scanner) errorToken(err string) Token {
	return NewToken(TOKEN_ERROR, err, int(len(err)), s.line)
}

func (s *Scanner) advance() rune {
	s.at++
	s.current = s.source[s.at]
	return s.source[s.at-1]
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	if s.current != expected {
		return false
	}
	s.at++
	s.current = s.source[s.at]
	return true
}

func (s *Scanner) skipWhiteSpace() {
	for {
		char := s.peek()
		switch char {
		case ' ':
		case '\r':
		case '\t':
			s.advance()
		case '\n':
			s.line++
			s.advance()
		case '/':
			if s.peekNext() == '/' {
				for s.peek() != '\n' && !s.isAtEnd() {
					s.advance()
				}
			} else {
				return
			}
		default:
			return
		}
	}
}

func (s *Scanner) isAlpha(char rune) bool {
	return unicode.IsLetter(char) || char == '_'
}

func (s *Scanner) identifier() Token {
	for s.isAlpha(s.peek()) || unicode.IsDigit(s.peek()) {
		s.advance()
	}
	return s.makeToken(s.identifierType())
}

func (s *Scanner) identifierType() TokenType {
	switch s.source[0] {
	case 'a':
		return s.checkKeyword(1, 2, "nd", TOKEN_AND)
	case 'c':
		return s.checkKeyword(1, 4, "lass", TOKEN_CLASS)
	case 'e':
		return s.checkKeyword(1, 3, "lse", TOKEN_ELSE)
	case 'f':
		{
			switch s.source[1] {
			case 'a':
				return s.checkKeyword(2, 3, "lse", TOKEN_FALSE)
			case 'o':
				return s.checkKeyword(2, 1, "r", TOKEN_FOR)
			case 'u':
				return s.checkKeyword(2, 1, "n", TOKEN_FUN)
			}
		}
	case 'i':
		return s.checkKeyword(1, 1, "f", TOKEN_IF)
	case 'n':
		return s.checkKeyword(1, 2, "il", TOKEN_NIL)
	case 'o':
		return s.checkKeyword(1, 1, "r", TOKEN_OR)
	case 'p':
		return s.checkKeyword(1, 4, "rint", TOKEN_PRINT)
	case 'r':
		return s.checkKeyword(1, 5, "eturn", TOKEN_RETURN)
	case 's':
		return s.checkKeyword(1, 4, "uper", TOKEN_SUPER)
	case 't':
		{
			switch s.source[1] {
			case 'h':
				return s.checkKeyword(2, 2, "is", TOKEN_THIS)
			case 'r':
				return s.checkKeyword(2, 2, "ue", TOKEN_TRUE)
			}
		}
	case 'v':
		return s.checkKeyword(1, 2, "ar", TOKEN_VAR)
	case 'w':
		return s.checkKeyword(1, 4, "hile", TOKEN_WHILE)

	}
	return TOKEN_IDENTIFIER
}

func (s *Scanner) checkKeyword(start, length int, rest string, tokenType TokenType) TokenType {
	pattern := "^" + rest + "$"
	re := regexp.MustCompile(pattern)

	keyword := string(s.source[start : start+length])

	if re.MatchString(keyword) {
		return tokenType
	}

	return TOKEN_IDENTIFIER
}

func (s *Scanner) number() Token {
	for unicode.IsDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && unicode.IsDigit(s.peekNext()) {
		s.advance()
		for unicode.IsDigit(s.peek()) {
			s.advance()
		}
	}

	return s.makeToken(TOKEN_NUMBER)
}

func (s *Scanner) str() Token {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		return s.errorToken("unterminated string")
	}
	s.advance()
	return s.makeToken(TOKEN_STRING)
}

func (s *Scanner) peek() rune {
	return s.current
}

func (s *Scanner) peekNext() rune {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current+1]
}
