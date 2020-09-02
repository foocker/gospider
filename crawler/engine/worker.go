package engine

import (
	"gospider/crawler/fetcher"

	log "github.com/sirupsen/logrus"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Error("request url %s failed", r.Url)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}
