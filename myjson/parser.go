package myjson

import "unicode/utf8"

type Parser struct {
	text   []rune
	curpos int
	len    int
}

func NewParser(text string) *Parser {
	return &Parser{
		text:   []rune(text),
		curpos: 0,
		len:    utf8.RuneCountInString(text),
	}
}

func (p *Parser) parseSymbol() (string, error) {
	now := p.curpos
	for now < p.len && p.text[now] == ' ' {
		now++
	}

	symbol := []rune{}
	for now < len(p.text) && p.text[now] != ' ' {
		symbol = append(symbol, p.text[now])
		now++
	}
	return string(symbol), nil
}

// {key: value} となるようなもの, json と区別
func (p *Parser) parseObject() (json map[string]interface{}, err error) {
	json = map[string]interface{}{}
	key, err := p.parseString()
	if err != nil {
		// パースエラー
		return
	}
	val, err := p.parseObject()
	if err != nil {
		// パースエラー
		return
	}
	json[key] = val
	// TODO 複数項目に対応
	return
}

// [value, value ,..]となるようなもの
func (p *Parser) parseList() {}

func (p *Parser) parseString() (string, error) { return "", nil }

func (p *Parser) parseNumber() {}

func (p *Parser) parseBool() (bool, error) {
	symbol, err := p.parseSymbol()
	if err != nil {
		return false, err
	}
	switch symbol {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		// パースエラー　got unknown symbol
		return false, err
	}
}

func (p *Parser) parseNull() (bool, error) {
	symbol, err := p.parseSymbol()
	if err != nil {
		return false, err
	}
	switch symbol {
	case "null":
		return true, nil
	default:
		return false, nil
	}
}

func (p *Parser) ParseJson() map[string]interface{} {
	p.parseObject()
	p.parseList()
	p.parseString()
	p.parseNumber()
	p.parseNull()
	p.parseBool()
	// どれにも当てはまらなければエラー
	return nil
}
