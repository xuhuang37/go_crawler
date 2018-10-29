package Persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"go_crawler/engine"
	"github.com/pkg/errors"
)

func ItemSaver() (chan engine.Item ,error){
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {

	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d, %v", itemCount, item)
			itemCount++
			err := save(item,client)
			if err != nil {
				log.Printf("Item Saver: error saving item %v, %v", item, err)
			}
		}
	}()
	return out,nil
}

func save(item engine.Item,client *elastic.Client) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).Id(item.Id).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil

}
