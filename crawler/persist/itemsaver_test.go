package persist

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic/v7"

	"gospider/crawler/clusterbaseinfo"
	"gospider/crawler/engine"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:     "https://book.douban.com/subject/4913064/",
		Type:    "douban",
		Id:      "4913064",
		Payload: clusterbaseinfo.BookExample,
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "book-test"
	// Save expected item
	err = Save(client, index, expected)

	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().Index(index).Type(expected.Type).
		Id(expected.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("Get from client: %s", resp.Source)

	var actual engine.Item
	json.Unmarshal(resp.Source, &actual)

	actualBookinfo, _ := clusterbaseinfo.FromJsonObj(actual.Payload)

	actual.Payload = actualBookinfo

	// Verfit result
	if actual != expected {
		t.Errorf("got %v; expected %v",
			actual, expected)
	}
}
