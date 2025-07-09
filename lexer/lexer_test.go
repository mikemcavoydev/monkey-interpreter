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

	!-/*5;
	5 < 10 > 5;

	if(5 < 10){
		return true;
	} else {
		return false; 
	}

	10 == 10;
	10 != 9;
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
			expectedType:    token.BANG,
			expectedLiteral: "!",
		},
		{
			expectedType:    token.MINUS,
			expectedLiteral: "-",
		},
		{
			expectedType:    token.SLASH,
			expectedLiteral: "/",
		},
		{
			expectedType:    token.ASTERIX,
			expectedLiteral: "*",
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
			expectedType:    token.INT,
			expectedLiteral: "5",
		},
		{
			expectedType:    token.LT,
			expectedLiteral: "<",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "10",
		},
		{
			expectedType:    token.GT,
			expectedLiteral: ">",
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
			expectedType:    token.IF,
			expectedLiteral: "if",
		},
		{
			expectedType:    token.LPAREN,
			expectedLiteral: "(",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "5",
		},
		{
			expectedType:    token.LT,
			expectedLiteral: "<",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "10",
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
			expectedType:    token.RETURN,
			expectedLiteral: "return",
		},
		{
			expectedType:    token.TRUE,
			expectedLiteral: "true",
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
			expectedType:    token.ELSE,
			expectedLiteral: "else",
		},
		{
			expectedType:    token.LBRACE,
			expectedLiteral: "{",
		},
		{
			expectedType:    token.RETURN,
			expectedLiteral: "return",
		},
		{
			expectedType:    token.FALSE,
			expectedLiteral: "false",
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
			expectedType:    token.INT,
			expectedLiteral: "10",
		},
		{
			expectedType:    token.EQ,
			expectedLiteral: "==",
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
			expectedType:    token.INT,
			expectedLiteral: "10",
		},
		{
			expectedType:    token.NOT_EQ,
			expectedLiteral: "!=",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "9",
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
