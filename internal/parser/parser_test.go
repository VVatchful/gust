package parser

import (
	"github.com/voidwyrm-2/gust/internal/lexer"
	"testing"
  "fmt"
)


func TestBasicParsing(t *testing.T) {
	input := `
let x = 5
let y = "hello"
let add = fn(a, b) { a + b }
if (x < 10) { x } else { y }
return add(x, 15)
!true
`
	t.Logf("Testing basic parsing with input:\n%s", input)

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 6 {
		t.Fatalf("program.Statements does not contain 6 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedType string
	}{
		{"*parser.LetStatement"},
		{"*parser.LetStatement"},
		{"*parser.LetStatement"},
		{"*parser.ExpressionStatement"},
		{"*parser.ReturnStatement"},
		{"*parser.ExpressionStatement"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		typeOf := stringify(stmt)
		t.Logf("Statement %d: expected=%q, got=%q", i, tt.expectedType, typeOf)
		if typeOf != tt.expectedType {
			t.Errorf("program.Statements[%d] type wrong. expected=%q, got=%q",
				i, tt.expectedType, typeOf)
		}
	}
}

func TestLetStatement(t *testing.T) {
	input := "let x = 5"
	t.Logf("Testing let statement parsing with input: %q", input)

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*LetStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not LetStatement. got=%T",
			program.Statements[0])
	}

	t.Logf("Let statement name: %s", stmt.Name.Value)
	if stmt.Name.Value != "x" {
		t.Errorf("letStmt.Name.Value not 'x'. got=%s", stmt.Name.Value)
	}

	if !testIntegerLiteral(t, stmt.Value, 5) {
		return
	}
	t.Logf("Let statement value correctly parsed as: %d", 5)
}

func TestFunctionLiteral(t *testing.T) {
	input := "fn(x, y) { x + y }"
	t.Logf("Testing function literal parsing with input: %q", input)

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ExpressionStatement. got=%T",
			program.Statements[0])
	}

	function, ok := stmt.Expression.(*FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not FunctionLiteral. got=%T", stmt.Expression)
	}

	t.Logf("Function parameters count: %d", len(function.Parameters))
	if len(function.Parameters) != 2 {
		t.Fatalf("function parameters wrong. want 2, got=%d",
			len(function.Parameters))
	}

	t.Logf("Function parameters: %s, %s", function.Parameters[0].Value, function.Parameters[1].Value)
	if function.Parameters[0].Value != "x" || function.Parameters[1].Value != "y" {
		t.Errorf("parameter values wrong. want 'x' and 'y', got=%q and %q",
			function.Parameters[0].Value, function.Parameters[1].Value)
	}

	t.Logf("Function body statements count: %d", len(function.Body.Statements))
	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has wrong number of statements. got=%d",
			len(function.Body.Statements))
	}
}

func TestIfExpression(t *testing.T) {
	input := `if (x < y) { x } else { y }`
	t.Logf("Testing if expression parsing with input: %q", input)

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ExpressionStatement. got=%T",
			program.Statements[0])
	}

	exp, ok := stmt.Expression.(*IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not IfExpression. got=%T", stmt.Expression)
	}

	t.Logf("If expression has alternative: %v", exp.Alternative != nil)
	if exp.Alternative == nil {
		t.Error("if expression has no alternative")
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for i, msg := range errors {
		t.Errorf("parser error %d: %q", i+1, msg)
	}
	t.FailNow()
}

func stringify(stmt Statement) string {
	switch stmt.(type) {
	case *LetStatement:
		return "*parser.LetStatement"
	case *ReturnStatement:
		return "*parser.ReturnStatement"
	case *ExpressionStatement:
		return "*parser.ExpressionStatement"
	default:
		return fmt.Sprintf("unknown: %T", stmt)
	}
}

func testIntegerLiteral(t *testing.T, il Expression, value int64) bool {
	integ, ok := il.(*IntegerLiteral)
	if !ok {
		t.Errorf("il not *IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	return true
}
