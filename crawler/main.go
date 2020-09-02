package main

import (
	"regexp"

	"gospider/crawler/config"
	"gospider/crawler/douban/parser/book"
	"gospider/crawler/engine"
	"gospider/crawler/persist"
	"gospider/crawler/scheduler"
)

var (
	starRe         = regexp.MustCompile(`"rating_per">([\d\.]+)%</span>`) // five results
	bookCategoryRe = regexp.MustCompile(`<a href="(/tag/[^"]+)">([^<]+)</a>`)
)

func main() {
	// Just test Worker
	//url := "https://book.douban.com/tag/?view=type&icn=index-sorttags-all"
	//body, _ := fetcher.Fetch(url)
	////fmt.Printf("%v", string(body))
	//match := bookCategoryRe.FindAllSubmatch(body, 32)
	//fmt.Printf("%v", len(match))
	////fmt.Println(string(match[4][1]))
	//for i, v := range match {
	//	fmt.Println(i, string(v[1]))
	//}
	//x := strings.Split(url, "/")
	//fmt.Println(x[len(x)-2], len(x))

	// Testing SimpleEngine
	//var seeds engine.Request
	//seeds = engine.Request{
	//	Url:    "https://book.douban.com/tag/?view=type&icn=index-sorttags-all",
	//	Parser: engine.NewFuncParser(book.ParseBookClass, "ParseBookClass")}
	//
	//e := engine.SimpleEngine{}
	//e.Run(seeds)

	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      4,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url: "https://book.douban.com/tag/?view=type&icn=index-sorttags-all",
		Parser: engine.NewFuncParser(
			book.ParseBookClass, config.BookClassParse),
	})
}
