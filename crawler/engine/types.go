package engine

import "gospider/crawler/config"

// func type
type ParseFunc func(body []byte, url string) ParseResult

// there are many page need diff parser
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// one url, one parser func
type Request struct {
	Url    string
	Parser Parser
}

// Item represent the parse result of one page, has many items
type ParseResult struct {
	Requests []Request
	Items    []Item
}

// one item should has its own struct, use interface{}
type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

// Consider Serialize when rpc for distribution
type FuncParser struct {
	parser ParseFunc
	name   string
}

func (p *FuncParser) Parse(contents []byte, url string) ParseResult {
	return p.parser(contents, url)
}
func (p *FuncParser) Serialize() (name string, args interface{}) {
	return p.name, nil
}

func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

// Creat an empty Parse for test
type NilParse struct{}

func (NilParse) Parse(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (NilParse) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}
