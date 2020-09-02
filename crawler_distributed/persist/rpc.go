package persist

import (
	"github.com/olivere/elastic/v7"
	"gospider/crawler/engine"
	"gospider/crawler/persist"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

// wrap up, call the persist Save. creat new struct.
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}
