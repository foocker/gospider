package book

import (
	"gospider/crawler/engine"
	"regexp"
)

// ([\u4e00-\u9fa5]+)
var bookRe = regexp.MustCompile(`href="(https://book.douban.com/subject/[\d]+/)" title="([^"]+)"`)

func ParseBook(contents []byte, _ string) engine.ParseResult {
	re := engine.ParseResult{}
	match := bookRe.FindAllSubmatch(contents, -1)
	for _, m := range match {
		re.Requests = append(re.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewParseBookInfo(string(m[2])),
		})
	}
	return re
}
