package persist

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/persist"
	"gopkg.in/olivere/elastic.v5"
)

type ItemService struct {
	Client *elastic.Client
}

func (s *ItemService) Save(item engine.Item, result *string) error {

	err := persist.Save(s.Client, item)

	if err == nil {
		*result = "OK"
	}

	return err

}
