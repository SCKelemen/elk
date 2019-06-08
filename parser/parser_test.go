package parser

import (
	"fmt"
	"testing"

	"github.com/SCKelemen/elk/ast"
	"github.com/SCKelemen/elk/scanner"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foo = 838383;
	`
	//input = fmt.Sprintf("%s%s", input, string(rune(0)))
	s := scanner.New(input)
	p := New(s)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() return nil")
	}
	for ind, st := range program.Statements {
		fmt.Printf("%v\t%s\n", ind, st)
		t.Logf("%v\t%s\n", ind, st)

	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements doesnt contain 3. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			t.Logf("%v %s", i, stmt)
			return
		}
	}
}
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not let, got=%q", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s, got %s", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("letstmt.Name.TokenLit npt %s, got %s", name, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}
