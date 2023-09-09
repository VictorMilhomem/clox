package compiler

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
	}

	return s.errorToken("unexpected character")
}

func (s *Scanner) isAtEnd() bool {
	return s.at == s.length
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

func (s *Scanner) peek() rune {
	return s.current
}

func (s *Scanner) peekNext() rune {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current+1]
}
