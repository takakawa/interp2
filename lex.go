package main

import (
	"io"
	"log"
	"strconv"
	"text/scanner"
)

type Lexer struct {
	s         scanner.Scanner
	program   []Statement
	hasErrors bool
}

func NewLexer(r io.Reader) *Lexer {
	l := new(Lexer)
	l.s.Init(r)
	l.s.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats | scanner.ScanStrings | scanner.SkipComments
	return l
}

var lexKeywords = map[string]int{
	"IF":    IF,
	"THEN":  THEN,
	"ELSE":  ELSE,
	"END":   END,
	"WHILE": WHILE,
	"DO":    DO,
	"PRINT": PRINT,
	"AND":   AND,
	"OR":    OR,
	"NOT":   NOT,
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok := l.s.Scan()
	switch tok {
	case scanner.EOF:
		return 0
	case scanner.Int, scanner.Float:
		lval.num, _ = strconv.ParseFloat(l.s.TokenText(), 64)
		return NUMBER
	case scanner.Ident:
		ident := l.s.TokenText()
		keyword, isKeyword := lexKeywords[ident]
		if isKeyword {
			return keyword
		}
		lval.str = ident
		return IDENTIFIER
	case scanner.String:
		text := l.s.TokenText()
		text = text[1 : len(text)-1]
		lval.str = text
		return STRING
	default:
		return int(tok)
	}
}

func (l *Lexer) Error(s string) {
	log.Println("Parser:", s)
	l.hasErrors = true
}

func (l *Lexer) Program() []Statement {
	if l.hasErrors {
		return nil
	}
	return l.program
}
