package parser

import (
	"testing"

	"github.com/SCKelemen/elk/ast"
	"github.com/SCKelemen/elk/scanner"
)

func TestValStatement(t *testing.T) {
	input := `
	val x = 5;
	val y = 10;
	val foobar = 838383;
	`
	s := scanner.New(input)
	p := New(s)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testValStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func testValStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "val" {
		t.Errorf("s.TokenLiteral not 'val'. got=%q", s.TokenLiteral())
		return false
	}

	valStmt, ok := s.(*ast.ValStatement)
	if !ok {
		t.Errorf("s not *ast.ValStatement. got=%T", s)
		return false
	}

	if valStmt.Name.Value != name {
		t.Errorf("valStmt.Name.Value not '%s'. got=%s", name, valStmt.Name.Value)
		return false
	}

	if valStmt.Name.TokenLiteral() != name {
		t.Errorf("valStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, valStmt.Name.TokenLiteral())
		return false
	}

	return true
}
