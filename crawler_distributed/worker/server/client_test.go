package main

import (
	"fmt"
	"gospider/crawler/config"
	rpcconfig "gospider/crawler_distributed/config"
	"gospider/crawler_distributed/rpcsupport"
	"gospider/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "https://book.douban.com/subject/4913064/",
		Parser: worker.SerializeParser{
			Name: config.BookInfoParse,
			Args: "活着",
		},
	}
	fmt.Println(req.Parser.Name, req.Parser.Args)
	var result worker.ParseResult
	err = client.Call(rpcconfig.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
