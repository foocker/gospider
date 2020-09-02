package config

const (
	// Parser names
	BookParse      = "ParseBook"
	BookInfoParse  = "ParseBookInfo"
	BookClassParse = "ParseBookClass"

	NilParser = "NilParser"

	// ElasticSearch
	ElasticIndex = "book_profile"
	ElasticType  = "douban"

	// Rate limiting
	Qps = 2
)
