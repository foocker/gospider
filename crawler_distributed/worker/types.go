package worker

import (
	"errors"
	"fmt"
	"gospider/crawler/config"
	"gospider/crawler/douban/parser/book"
	"gospider/crawler/engine"
	"log"
)

// wrap up, new Parser should has its name for Serialize(using in rpc)
type SerializeParser struct {
	Name string
	Args interface{} // diff book has diff name(for other func which has diff parameters)
}

// new Serialize Parser, just some name info, but connect to original Request in config
type Request struct {
	Url    string
	Parser SerializeParser
}

// new Parser Result
type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

// Request Serialize
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializeParser{
			Name: name,
			Args: args,
		},
	}
}

// Result Serialize
func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		ReqDe, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error when desialize request %v, and error is %v", req, err)
			continue
		}
		result.Requests = append(result.Requests, ReqDe)
	}
	return result
}

func deserializeParser(p SerializeParser) (engine.Parser, error) {
	switch p.Name {
	case config.BookClassParse:
		return engine.NewFuncParser(book.ParseBookClass, config.BookClassParse), nil
	case config.BookParse:
		return engine.NewFuncParser(book.ParseBook, config.BookParse), nil
	case config.BookInfoParse:
		if userName, ok := p.Args.(string); ok {
			return book.NewParseBookInfo(userName), nil
		} else {
			return nil, fmt.Errorf("invalid"+"arg: %v", p.Args)
		}
	case config.NilParser:
		return engine.NilParse{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}
