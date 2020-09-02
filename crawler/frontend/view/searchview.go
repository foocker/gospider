package view

import (
	"html/template"

	"gospider/crawler/frontend/page"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(fname string) SearchResultView{
	return SearchResultView{
		template:template.Must(template.ParseFiles(fname))
	}
}

func (s SearchResultView) Render(w io.Writer, data page.SearchResult) error{
	return s.template.Execute(w, data)
}

