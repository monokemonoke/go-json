package myjson

import (
	"fmt"
)

type Lexer struct {
	source string
	curpos int
}

func NewLexer(source string) *Lexer {
	var lexer Lexer
	lexer.source = source
	lexer.curpos = 0
	return &lexer
}

func (l *Lexer) getCurChar() (string, bool) {
	if len(l.source) <= l.curpos {
		return "", false
	}
	return string(l.source[l.curpos]), true
}

func (l *Lexer) getCurCharAndSkip() (string, bool) {
	if len(l.source) <= l.curpos {
		return "", false
	}
	char := string(l.source[l.curpos])
	l.curpos++
	return char, true
}

func (l *Lexer) skipSpaces() {
	for {
		char, ok := l.getCurChar()
		if !ok || char != " " {
			return
		}
		l.curpos++
	}
}

func (l *Lexer) getStringToken() (name string) {
	l.skipSpaces()
	char, ok := l.getCurChar()
	if !ok || char != "\"" {
		fmt.Println("early return")
		fmt.Println(char, ok)
		return
	}
	l.curpos++

	for {
		char, ok := l.getCurChar()
		fmt.Println(char, ok)
		if !ok || char == "\"" {
			l.curpos++
			return
		}
		name += char
		l.curpos++
	}
}

func (l *Lexer) Lex() {
	l.skipSpaces()
	l.getCurCharAndSkip() // -> {
	key := l.getStringToken()
	l.skipSpaces()
	l.getCurCharAndSkip() // -> :
	val := l.getStringToken()
	l.skipSpaces()
	l.getCurCharAndSkip() // -> }
	fmt.Println(key, val)
}
