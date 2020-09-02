package page

// Show on page
type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PreForm  int
	NextForm int
	Items    []interface{}
}
