package persist

import (
	"context"
	"crawl/engine"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

//func ItemSave() chan interface{ }{
//
//	out:=make(chan  interface{})
//
//	go func() {
//		itemcount:=0
//
//		for {
//			item:=<-out
//			 log.Printf("Item Saver : Got%d,%v",itemcount,item )
//
//			itemcount++
//		}
//	}()
//
//	return out
//}
func ItemSave() chan engine.Item {

	out := make(chan engine.Item)

	go func() {
		itemcount := 0

		for {
			item := <-out
			log.Printf("Item saver:Got$%d,%v", itemcount, item)
			save(item)
			itemcount++
		}

	}()

	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().Index("dating_profile").Type(item.Type).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)

	}
	_, err = indexService.Do(context.Background())

	if err != nil {
		panic(err)
	}

	return nil

}
