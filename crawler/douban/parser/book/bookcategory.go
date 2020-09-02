package book

import (
	"gospider/crawler/engine"
	"regexp"
)

var (
	bookCategoryRe = regexp.MustCompile(`<a href="(/tag/[^"]+)">([^<]+)</a>`)
)

func ParseBookClass(contents []byte, _ string) engine.ParseResult {
	re := engine.ParseResult{}

	match := bookCategoryRe.FindAllSubmatch(contents, -1)
	for _, m := range match {
		re.Requests = append(re.Requests, engine.Request{
			Url:    "https://book.douban.com/" + string(m[1]),
			Parser: engine.NewFuncParser(ParseBook, "ParseBook"),
		})
	}
	return re
}
