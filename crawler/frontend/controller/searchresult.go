package controller

import (
	"github.com/olivere/elastic/v7"

	"gospider/crawler/frontend/view"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {

}

func (h SearchResultHandler) ServeHTTP() {

}

const pagesize = 10

func (h SearchResultHandler) getSearchResult() {

}

func rewriteQuerySring(q string) string {

}
