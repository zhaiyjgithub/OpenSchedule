package database

import (
	"OpenSchedule/src/conf"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
)

const DriverName = "mysql"

var (
	elasticSearchOnce sync.Once
	elasticSearchEngine *elastic.Client

	mysqlOnce sync.Once
	mysqlEngine *gorm.DB

	localMySqlOnce sync.Once
	localMySqlEngine *gorm.DB
)

func GetMySqlEngine() *gorm.DB  {
	mysqlOnce.Do(func() {
		var err error
		c := conf.MySQLConf
		driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
			c.User, c.Password, c.Host, c.Port, c.DatabaseName)
		localMySqlEngine, err = gorm.Open(mysql.Open(driveSource), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatal(err)
		}
	})

	return localMySqlEngine
}

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
