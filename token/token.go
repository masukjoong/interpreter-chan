package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// TokenType
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 식별자 + 리터럴
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// 연산자
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// 구분자
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// 예약어
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// keyword literal
const (
	FUNCTION_LITERAL = "fn"
	LET_LITERAL      = "let"
	TRUE_LITERAL     = "true"
	FALSE_LITERAL    = "false"
	IF_LITERAL       = "if"
	ELSE_LITERAL     = "else"
	RETURN_LITERAL   = "return"
)

var keywords = map[string]TokenType{
	FUNCTION_LITERAL: FUNCTION,
	LET_LITERAL:      LET,
	TRUE_LITERAL:     TRUE,
	FALSE_LITERAL:    FALSE,
	IF_LITERAL:       IF,
	ELSE_LITERAL:     ELSE,
	RETURN_LITERAL:   RETURN,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
