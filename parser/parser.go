package parser

import (
	"fmt"

	"github.com/SCKelemen/elk/ast"
	"github.com/SCKelemen/elk/scanner"
	"github.com/SCKelemen/elk/token"
)

type Parser struct {
	s       *scanner.Scanner
	current token.Token
	next    token.Token
	errata  []string
}

func New(s *scanner.Scanner) *Parser {
	p := &Parser{s: s, errata: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errata
}

func (p *Parser) nextToken() {
	p.current = p.next
	p.next = p.s.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := &ast.Program{}
	prog.Statements = []ast.Statement{}

	for p.current.Kind != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		p.nextToken()
	}

	return prog
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.current.Kind {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {

	stmt := &ast.LetStatement{Token: p.current}

	if !p.expectNext(token.IDENTITY) {
		//fmt.Printf("expected to be followed by IDENTITY, but was followed by %T \n", next)
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.current, Value: p.current.Literal}

	if !p.expectNext(token.EQL) {
		return nil
	}

	// skip exp
	for !p.isCurrent(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) isCurrent(kind token.TokenKind) bool {
	return p.current.Kind == kind
}

func (p *Parser) isNext(kind token.TokenKind) bool {
	return p.next.Kind == kind
}
func (p *Parser) expectNext(kind token.TokenKind) bool {
	if p.isNext(kind) {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) peekError(t token.TokenKind) {
	msg := fmt.Sprintf("expected next token to be of type %s, got %s instead", t, p.next.Kind)
	p.errata = append(p.errata, msg)
}
