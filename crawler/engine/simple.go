package engine

import "log"

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var itemCount int
	var requestsQue []Request
	for _, r := range seeds {
		requestsQue = append(requestsQue, r)
	}
	for len(requestsQue) > 0 {
		r := requestsQue[0]
		requestsQue = requestsQue[1:]

		parseresult, err := Worker(r)
		if err != nil {
			continue
		}
		requestsQue = append(requestsQue, parseresult.Requests...)
		for _, item := range parseresult.Items {
			itemCount++
			log.Printf("Got item %d %v", itemCount, item)
		}
	}
}
