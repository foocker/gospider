package book

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"

	"gospider/crawler/clusterbaseinfo"
	"gospider/crawler/config"
	"gospider/crawler/engine"
)

// use some other way to get the struct info, not regexp

var (
	// book struct
	urlImageRe = regexp.MustCompile(`href="https://img3.doubanio.com/view/subject/l/public/[\w]+.jpg" title="([^>]+)">`)
	//booknameRe   from title
	writerRe          = regexp.MustCompile(`<span class="pl">作者</span>[^<]+<a class=([^>]+>([^<]+)</a>)*]`)
	publisherRe       = regexp.MustCompile(`<span class="pl">出版社:</span> ([^<]+)<br/>`)
	publicationYearRe = regexp.MustCompile(`<span class="pl">出版年:</span> ([^<]+)<br/>`)
	numPagesRe        = regexp.MustCompile(`<span class="pl">页数:</span> ([\d]+)<br/>`)
	pricingRe         = regexp.MustCompile(`<span class="pl">定价:</span> ([\d\.]+)[\s\S]?<br/>`)
	//finishingRe       = regexp.MustCompile(`<span class="pl">装帧:</span> ([\u4e00-\u9fa5]{2,})<br/>`) // some no
	isbnRe = regexp.MustCompile(`<span class="pl">ISBN:</span> ([\d]+)<br/>`)
	// star struct
	scoresRe      = regexp.MustCompile(`property="v:average"> ([\d\.]+) </strong>`)
	numconmentsRe = regexp.MustCompile(`"v:votes">([\d]+)</span>人评价`)
	starRe        = regexp.MustCompile(`"rating_per">([\d\.]+)%</span>`) // five results
)

func bs4(html []byte) {
	doc := soup.HTMLParse(string(html))
	zuozhe := doc.Find("div", "id", "info").Find("span").FindAll("a")
	for _, v := range zuozhe {
		fmt.Println(v.Text())
	}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) == 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func ParseBookInfo(contents []byte, url string) engine.ParseResult {
	re := engine.ParseResult{}
	//bookinfo := clusterbaseinfo.BookExample
	bookinfo := clusterbaseinfo.BookInfo{}
	bookinfo.Url = url
	bookinfo.PublicationYear = extractString(contents, publicationYearRe)
	bookinfo.Publisher = extractString(contents, publisherRe)
	bookinfo.Writer = extractString(contents, writerRe)
	bookinfo.Name = extractString(contents, urlImageRe)

	rampage, err := strconv.Atoi(extractString(contents, numPagesRe))
	if err == nil {
		bookinfo.NumPages = int64(rampage)
	}
	isbn, err := strconv.Atoi(extractString(contents, isbnRe))
	if err == nil {
		bookinfo.ISBN = int64(isbn)
	}
	price, err := strconv.ParseFloat(extractString(contents, pricingRe), 64)
	if err == nil {
		bookinfo.Pricing = price
	}

	sarina := clusterbaseinfo.Star{}
	score, err := strconv.ParseFloat(extractString(contents, scoresRe), 64)
	if err == nil {
		sarina.Scores = score
	}
	newcomers, err := strconv.Atoi(extractString(contents, numconmentsRe))
	if err == nil {
		sarina.NumComments = int64(newcomers)
	}
	starch := starRe.FindAllSubmatch(contents, 32)
	for i, v := range starch {
		sarina.Stars[i], _ = strconv.ParseFloat(string(v[1]), 32) // err not exist always?
	}
	bookinfo.Starinfo = sarina

	item := engine.Item{}
	item.Url = url
	x := strings.Split(url, "/")
	item.Id = x[len(x)-2]
	item.Payload = bookinfo
	item.Type = config.ElasticType

	re.Items = append(re.Items, item)

	return re
}

// just use bookname to complete the Parser interface
type BookInfoParser struct {
	bookName string // every book has diff name
}

//
func (b *BookInfoParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseBookInfo(contents, url)
}

func (b *BookInfoParser) Serialize() (name string, args interface{}) {
	return b.bookName, nil
}

func NewParseBookInfo(name string) *BookInfoParser {
	return &BookInfoParser{bookName: name}
}
