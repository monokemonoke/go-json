package myjson

import (
	"log"
)

type Lexer struct {
	source string
	curpos int
}

type Json map[string]interface{}

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
	l.skipSpaces()
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
		return
	}
	l.curpos++

	for {
		char, ok := l.getCurChar()
		if !ok || char == "\"" {
			l.curpos++
			return
		}
		name += char
		l.curpos++
	}
}

func (l *Lexer) expectChar(char string) {
	c, ok := l.getCurCharAndSkip() // -> {
	if !ok || c != char {
		log.Printf("Expect %s got %s\n", char, c)
		return
	}
}

func (l *Lexer) Lex() Json {
	json := Json{}
	l.expectChar("{")

	for {
		key := l.getStringToken()
		log.Println(key)
		l.expectChar(":")
		val := l.getStringToken()
		log.Println(val)
		json[key] = val

		c, ok := l.getCurCharAndSkip()
		if !ok || c != "," {
			return json
		}
	}
}
