package database

import (
	"OpenSchedule/src/conf"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"sync"
)

var (
	elasticSearchOnce sync.Once
	elasticSearchEngine *elastic.Client

)

func GetElasticSearchEngine() *elastic.Client {
	elasticSearchOnce.Do(func() {
		var err error
		url := fmt.Sprintf("%s:%d", conf.ElasticSearchConf.Host, conf.ElasticSearchConf.Port)
		elasticSearchEngine, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(url))
		if err != nil {
			log.Fatalf("Setup elastic search server failed....: %v", err)
		}

		info, code, err := elasticSearchEngine.Ping(url).Do(context.Background())
		if err != nil {
			log.Fatalf("ping elastic search server error: %v", err)
		}

		fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

		version, err := elasticSearchEngine.ElasticsearchVersion(url)
		if err != nil {
			log.Fatalf("Get elastic version error: %v", err)
		}
		fmt.Printf("Elasticsearch version %s\n", version)
	})

	return elasticSearchEngine
}
