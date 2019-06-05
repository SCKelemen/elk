package scanner

import (
	"github.com/SCKelemen/elk/token"
	"github.com/SCKelemen/elk/util"
)

type Scanner struct {
	input   string
	head    int // location of pointer head
	read    int // read ahead pointer
	current rune
}

func New(input string) *Scanner {
	s := &Scanner{input: input}
	// load the first rune into the current
	// character under inspection for lexical
	// analysis
	s.readChar()
	return s
}

// readChar 's only responsibility is to progress
// the read-ahead head, check for EOF, and then
// update head to read-ahead head
func (s *Scanner) readChar() {
	// if the look-ahead pointer reaches
	// the end of the input stream,
	// set the current character to NUL/0
	// indicating EOF
	if s.read >= len(s.input) {
		s.current = 0
	} else {
		// else, set the current character
		// under inspection to be the char
		// at the look-ahead position
		s.current = rune(s.input[s.read])
	}

	// then we can set the head to the
	// read-ahead head
	s.head = s.read
	// and then increment the read-ahead head
	s.read++
}

func (s *Scanner) NextToken() token.Token {
	var tok token.Token

	s.skipWhitespace()

	switch s.current {
	case '.':
		tok = s.readDots()
	case '_':
		tok = newToken(token.UNDERSCORE, s.current)
	case ':':
		tok = newToken(token.COLON, s.current)
	case ';':
		tok = newToken(token.SEMICOLON, s.current)
	case '(':
		tok = newToken(token.LPAREN, s.current)
	case ')':
		tok = newToken(token.RPAREN, s.current)
	case '{':
		tok = newToken(token.LBRACE, s.current)
	case '}':
		tok = newToken(token.RBRACE, s.current)
	case '[':
		tok = newToken(token.LBRACK, s.current)
	case ']':
		tok = newToken(token.RBRACK, s.current)
	case '?':
		tok = newToken(token.EROTEME, s.current)
	case '!':
		tok = newToken(token.BANG, s.current)
	case 0:
		tok.Literal = ""
		tok.Kind = token.EOF
	default:
		if util.IsIdentifierInitialChar(s.current) {
			tok.Literal = s.readIdentifier()
			tok.Kind = token.Lookup(tok.Literal)
			return tok
		} else if util.IsNumericInitialChar(s.current) {
			tok.Kind = token.INTEGER
			tok.Literal = s.readNumber()
		} else if util.IsQuote(s.current) {
			tok.Kind = token.STRING
			tok.Literal = s.readStringLiteral()
		} else {
			tok = newToken(token.ILLEGAL, s.current)
		}
	}

	s.readChar()
	return tok
}

// readIdentifier
// Identifiers begin with '_' || Letter
// Identifiers may contain '_' || Letter || Digit
func (s *Scanner) readIdentifier() string {
	position := s.head
	for util.IsIdentifierChar(s.current) {
		s.readChar()
	}
	return s.input[position:s.head]
}

// readNumber
// Numbers begin with Digit
// Numbers may contain Digit || '_'
func (s *Scanner) readNumber() string {
	position := s.head
	for util.IsNumericChar(s.current) {
		s.readChar()
	}
	return s.input[position:s.head]
}

// readStringLiteral
//
//
func (s *Scanner) readStringLiteral() string {
	position := s.head
	s.read = s.head + 1

	for !util.IsQuote(rune(s.input[s.read])) {
		s.read++
	}
	s.read++ // need to read that last one
	s.head = s.read
	return s.input[position:s.head]
}

// readDots
// DOT: .
// DOTDOT: ..
// ELIPSIS: ...
func (s *Scanner) readDots() token.Token {
	position := s.head
	var dotCount = 1
	s.read = s.head + 1
	for rune(s.input[s.read]) == '.' && dotCount < 3 {
		dotCount++
		s.read++
	}
	s.head = s.read
	var kind = token.ILLEGAL
	switch dotCount {
	case 1:
		kind = token.DOT
	case 2:
		kind = token.DOTDOT
	case 3:
		kind = token.ELIPSIS
	}
	return token.Token{Kind: kind, Literal: s.input[position:s.head]}
}

// skipWhitespace 's only responsibility is to
// read while the current token under inspection
// remains a whitespace character. These don't have
// syntactic or semantic meaning to the language.
func (s *Scanner) skipWhitespace() {
	for util.IsWhitespace(s.current) {
		s.readChar()
	}
}

func newToken(kind token.TokenKind, ch rune) token.Token {
	return token.Token{Kind: kind, Literal: string(ch)}
}
