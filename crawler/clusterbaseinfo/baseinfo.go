package clusterbaseinfo

import "encoding/json"

type BookInfo struct {
	Publisher, PublicationYear string
	Writer, Name, Url          string
	NumPages, ISBN             int64
	Pricing                    float64
	Starinfo                   Star
}

type Star struct {
	NumComments int64
	Scores      float64
	Stars       [5]float64
}

func Map2BookInfo(x interface{}) BookInfo {
	str, _ := json.Marshal(x)
	y := BookInfo{}
	json.Unmarshal(str, &y)
	return y
}

// check save info is correct
func FromJsonObj(o interface{}) (BookInfo, error) {
	var bookinfo BookInfo
	s, err := json.Marshal(o)
	if err != nil {
		return bookinfo, err
	}
	err = json.Unmarshal(s, &bookinfo)
	return bookinfo, err
}

//var x = [5]float64{0.1, 0.1, 2.2, 22.2, 75.4}
var BookExample = BookInfo{
	Url:             "https://book.douban.com/subject/4913064/",
	Name:            "活着",
	Writer:          "余华",
	Publisher:       "作家出版社",
	PublicationYear: "2012-8-1",
	NumPages:        191,
	Pricing:         20.00,
	ISBN:            9787506365437,
	Starinfo: Star{NumComments: 519645,
		Scores: 9.4,
		Stars:  [5]float64{0.1, 0.1, 2.2, 22.1, 75.5}},
}
