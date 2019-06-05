package token

import "strconv"

type TokenKind int

type Token struct {
	Kind    TokenKind
	Literal string
}

const (
	ILLEGAL TokenKind = iota
	EOF
	TRIVIA

	IDENTITY

	//primitives
	literals
	INTEGER
	STRING
	literals_end
	// primitives_end

	punctuation
	LPAREN // (
	RPAREN // )

	LBRACE // {
	RBRACE // }

	// QUO        // "

	COLON      // :
	UNDERSCORE // _
	punctuation_end

	_keywords
	FUNC
	MATCH
	keywords_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	TRIVIA:  "TRIVIA",

	IDENTITY: "IDENTITY",

	INTEGER: "INTEGER",
	STRING:  "STRING",

	LPAREN: "(",
	RPAREN: ")",

	LBRACE: "{",
	RBRACE: "}",

	COLON:      ":",
	UNDERSCORE: "_",

	FUNC:  "func",
	MATCH: "match",
}

func (token TokenKind) String() string {
	s := ""
	// if the tokenkind is in the above list,
	// get the name
	if 0 <= token && token < TokenKind(len(tokens)) {
		s = tokens[token]
	}
	if s == "" {
		// if we couldn't get the name, make one up
		s = "token(" + strconv.Itoa(int(token)) + ")"
	}
	return s
}

var keywords map[string]TokenKind

func init() {
	keywords = make(map[string]TokenKind)
	for i := _keywords + 1; i < keywords_end; i++ {
		keywords[tokens[i]] = i
	}
}

func Lookup(ident string) TokenKind {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}
	return IDENTITY
}

var kw = map[string]TokenKind{
	"func":  FUNC,
	"match": MATCH,
}

func LookupIdent(ident string) TokenKind {
	if tok, ok := kw[ident]; ok {
		return tok
	}
	return IDENTITY
}

// goal: Lex the following:
/*

	func itoa(index: int): string {
		index match {
			1:	"one"
			2:	"two"
			3:	"three"
			4:	"four"
			_:	"Error"
		}
	}

*/

// result should look like:
/*
	{
		Token{Kind: FUNC, 		Literal: "func"},		func
		Token{Kind: IDENTITY, 	Literal: "itoa"},		itoa
		Token{Kind: LPAREN, 	Literal: "("},			(
		Token{Kind: IDENTITY, 	Literal: "index"},		index
		Token{Kind: COLON, 		Literal: ":"},			:
		Token{Kind: IDENTITY, 	Literal: "int"},		int
		Token{Kind: RPAREN, 	Literal: ")"},			)
		Token{Kind: COLON, 		Literal: ":"},			:
		Token{Kind: IDENTITY, 	Literal: "string"},		string
		Token{Kind: LBRACE, 	Literal: "{"},			{
		Token{Kind: IDENTITY, 	Literal: "index"},		index
		Token{Kind: MATCH, 		Literal: "match"},		match
		Token{Kind: LBRACE, 	Literal: "{"},			{
		Token{Kind: INTEGER, 	Literal: "1"},			1
		Token{Kind: COLON, 		Literal: ":"},			:
		Token{Kind: STRING, 	Literal: "\"one\""},	"one"
		Token{Kind: INTEGER, 	Literal: "2"},			2
		Token{Kind: COLON, 		Literal: ":"},			:
		Token{Kind: STRING, 	Literal: "\"two\""},	"two"
		Token{Kind: INTEGER, 	Literal: "3"},			3
		Token{Kind: COLON, 		Literal: ":"},			:
		Token{Kind: STRING, 	Literal: "\"three\""},	"three"
		Token{Kind: INTEGER, 	Literal: "4"},			4
		Token{Kind: COLON, 		Literal: ":"},			:
		Token{Kind: STRING, 	Literal: "\"four\""},	"four"
		Token{Kind: UNDERSCORE, Literal: "_"},			_
		Token{Kind: COLON, 		Literal: ":"},			:
		Token{Kind: STRING, 	Literal: "\"Error\""},	"Error"
		Token{Kind: RBRACE, 	Literal: "}"},			}
		Token{Kind: RBRACE, 	Literal: "}"},			}
	}
*/
