package main

import (
	"flag"
	"gospider/crawler/config"
	"gospider/crawler_distributed/persist"
	"gospider/crawler_distributed/rpcsupport"

	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen to")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serverRpc(
		fmt.Sprintf("%d", *port),
		config.ElasticIndex))

}

func serverRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
