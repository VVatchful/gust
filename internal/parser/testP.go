package parser

import (
	"github.com/voidwyrm-2/gust/internal/lexer"
	"testing"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLiteralExpression(t *testing.T, exp Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
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

func testIdentifier(t *testing.T, exp Expression, value string) bool {
	ident, ok := exp.(*Identifier)
	if !ok {
		t.Errorf("exp not *Identifier. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	return true
}

func testBooleanLiteral(t *testing.T, exp Expression, value bool) bool {
	bo, ok := exp.(*Boolean)
	if !ok {
		t.Errorf("exp not *Boolean. got=%T", exp)
		return false
	}
	if bo.Value != value {
		t.Errorf("bo.Value not %t. got=%t", value, bo.Value)
		return false
	}
	return true
}

// TestBasicParsing tests all basic expressions and statements in one go
func TestBasicParsing(t *testing.T) {
	input := `
let x = 5;
let y = true;
return 10;
!false;
-5;
5 + 10;
if (x < y) { x }
fn(x, y) { x + y; }
add(1, 2);
"hello";
`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 10 {
		t.Fatalf("program.Statements does not contain 10 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedType string
	}{
		{"*parser.LetStatement"},        // let x = 5
		{"*parser.LetStatement"},        // let y = true
		{"*parser.ReturnStatement"},     // return 10
		{"*parser.ExpressionStatement"}, // !false
		{"*parser.ExpressionStatement"}, // -5
		{"*parser.ExpressionStatement"}, // 5 + 10
		{"*parser.ExpressionStatement"}, // if expression
		{"*parser.ExpressionStatement"}, // fn expression
		{"*parser.ExpressionStatement"}, // call expression
		{"*parser.ExpressionStatement"}, // string literal
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if typeOf := TypeOf(stmt); typeOf != tt.expectedType {
			t.Errorf("stmt[%d] is not %s. got=%s",
				i, tt.expectedType, typeOf)
		}
	}
}

// Helper function to get type as string
func TypeOf(stmt Statement) string {
	switch stmt.(type) {
	case *LetStatement:
		return "*parser.LetStatement"
	case *ReturnStatement:
		return "*parser.ReturnStatement"
	case *ExpressionStatement:
		return "*parser.ExpressionStatement"
	default:
		return "unknown"
	}
}

// TestOperatorPrecedence verifies that operator precedence is correctly handled
func TestOperatorPrecedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"a + b * c",
			"(a + (b * c))",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b || c * d",
			"((a + b) || (c * d))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String() // You'll need to implement String() methods
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

// TestErrorHandling verifies that the parser properly handles syntax errors
func TestErrorHandling(t *testing.T) {
	input := `
let = 5;
let 123;
fn(x,) {};
`

	l := lexer.New(input)
	p := New(l)
	p.ParseProgram()

	errors := p.Errors()
	if len(errors) == 0 {
		t.Error("parser didn't report any errors for invalid input")
	}
}

// Implementation of String() methods for testing
func (p *Program) String() string {
	if len(p.Statements) > 0 {
		if es, ok := p.Statements[0].(*ExpressionStatement); ok {
			if ie, ok := es.Expression.(*InfixExpression); ok {
				return "(" + ie.Left.TokenLiteral() + " " + ie.Operator + " " + ie.Right.TokenLiteral() + ")"
			}
			if pe, ok := es.Expression.(*PrefixExpression); ok {
				return "(" + pe.Operator + pe.Right.TokenLiteral() + ")"
			}
		}
	}
	return ""
}
