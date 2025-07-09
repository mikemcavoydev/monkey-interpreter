package lexer

import (
	"testing"

	"github.com/mikemcavoydev/monkey-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};
	
	let result = add(five, ten);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{
			expectedType:    token.LET,
			expectedLiteral: "let",
		},
		{
			expectedType:    token.INDENTIFIER,
			expectedLiteral: "five",
		},
		{
			expectedType:    token.ASSIGN,
			expectedLiteral: "=",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "5",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},
		{
			expectedType:    token.LET,
			expectedLiteral: "let",
		},
		{
			expectedType:    token.INDENTIFIER,
			expectedLiteral: "ten",
		},
		{
			expectedType:    token.ASSIGN,
			expectedLiteral: "=",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "10",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},
		{
			expectedType:    token.LET,
			expectedLiteral: "let",
		},
		{
			expectedType:    token.INDENTIFIER,
			expectedLiteral: "add",
		},
		{
			expectedType:    token.ASSIGN,
			expectedLiteral: "=",
		},
		{
			expectedType:    token.FUNCTION,
			expectedLiteral: "fn",
		}, {
			expectedType:    token.LPAREN,
			expectedLiteral: "(",
		},
		{
			expectedType:    token.INDENTIFIER,
			expectedLiteral: "x",
		},
		{
			expectedType:    token.COMMA,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.INDENTIFIER,
			expectedLiteral: "y",
		},
		{
			expectedType:    token.RPAREN,
			expectedLiteral: ")",
		},
		{
			expectedType:    token.LBRACE,
			expectedLiteral: "{",
		},
		{
			expectedType:    token.INDENTIFIER,
			expectedLiteral: "x",
		},
		{
			expectedType:    token.PLUS,
			expectedLiteral: "+",
		},
		{
			expectedType:    token.INDENTIFIER,
			expectedLiteral: "y",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},
		{
			expectedType:    token.RBRACE,
			expectedLiteral: "}",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},
		{
			expectedType:    token.LET,
			expectedLiteral: "let",
		},
		{
			expectedType:    token.INDENTIFIER,
			expectedLiteral: "result",
		},
		{
			expectedType:    token.ASSIGN,
			expectedLiteral: "=",
		},
		{
			expectedType:    token.INDENTIFIER,
			expectedLiteral: "add",
		},
		{
			expectedType:    token.LPAREN,
			expectedLiteral: "(",
		},
		{
			expectedType:    token.INDENTIFIER,
			expectedLiteral: "five",
		},
		{
			expectedType:    token.COMMA,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.INDENTIFIER,
			expectedLiteral: "ten",
		},
		{
			expectedType:    token.RPAREN,
			expectedLiteral: ")",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},
		{
			expectedType:    token.EOF,
			expectedLiteral: "",
		},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i,
				tt.expectedType,
				tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}
