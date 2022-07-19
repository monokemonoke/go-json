package myjson

import (
	"log"
)

type Parser struct {
	source string
	curpos int
}

type Json map[string]interface{}

func NewParser(source string) *Parser {
	var newParser Parser
	newParser.source = source
	newParser.curpos = 0
	return &newParser
}

func (p *Parser) getCurChar() (string, bool) {
	if len(p.source) <= p.curpos {
		return "", false
	}
	return string(p.source[p.curpos]), true
}

func (p *Parser) getCurCharAndSkip() (string, bool) {
	p.skipSpaces()
	if len(p.source) <= p.curpos {
		return "", false
	}
	char := string(p.source[p.curpos])
	p.curpos++
	return char, true
}

func (p *Parser) skipSpaces() {
	for {
		char, ok := p.getCurChar()
		if !ok || char != " " {
			return
		}
		p.curpos++
	}
}

func (p *Parser) getStringToken() (name string) {
	p.skipSpaces()
	char, ok := p.getCurChar()
	if !ok || char != "\"" {
		return
	}
	p.curpos++

	for {
		char, ok := p.getCurChar()
		if !ok || char == "\"" {
			p.curpos++
			return
		}
		name += char
		p.curpos++
	}
}

func (p *Parser) expectChar(char string) {
	c, ok := p.getCurCharAndSkip() // -> {
	if !ok || c != char {
		log.Printf("Expect %s got %s\n", char, c)
		return
	}
}

func (p *Parser) Parse() Json {
	json := Json{}
	p.expectChar("{")

	for {
		key := p.getStringToken()
		log.Println(key)
		p.expectChar(":")
		val := p.getStringToken()
		log.Println(val)
		json[key] = val

		c, ok := p.getCurCharAndSkip()
		if !ok || c != "," {
			return json
		}
	}
}
