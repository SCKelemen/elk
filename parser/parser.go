package parser

import (
	"github.com/SCKelemen/elk/ast"
	"github.com/SCKelemen/elk/scanner"
	"github.com/SCKelemen/elk/token"
)

type Parser struct {
	s *scanner.Scanner

	ct token.Token
	nt token.Token
}

func New(s *scanner.Scanner) *Parser {
	p := &Parser{s: s}
	p.next()
	p.next()
	return p
}

func (p *Parser) next() {
	p.ct = p.nt
	p.nt = p.s.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.ct.Kind != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.next()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.ct.Kind {
	case token.VAL:
		return p.parseValStatement()
	default:
		return nil
	}
}

func (p *Parser) parseValStatement() *ast.ValStatement {
	stmt := &ast.ValStatement{Token: p.ct}

	if !p.expectNT(token.IDENTITY) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.ct, Value: p.ct.Literal}

	if !p.expectNT(token.EQL) {
		return nil
	}

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.isCT(token.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *Parser) isCT(kind token.TokenKind) bool {
	return p.ct.Kind == kind
}
func (p *Parser) isNT(kind token.TokenKind) bool {
	return p.nt.Kind == kind
}

func (p *Parser) expectNT(kind token.TokenKind) bool {
	if p.isNT(kind) {
		p.next()
		return true
	} else {
		return false
	}
}
